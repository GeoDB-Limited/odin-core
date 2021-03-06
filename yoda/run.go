package yoda

import (
	"context"
	"github.com/GeoDB-Limited/odin-core/yoda/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"path/filepath"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/log"
	httpclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/GeoDB-Limited/odin-core/pkg/filecache"
	oracletypes "github.com/GeoDB-Limited/odin-core/x/oracle/types"
	"github.com/GeoDB-Limited/odin-core/yoda/executor"
)

const (
	// TODO: We can subscribe only for txs that contain request messages
	TxQuery = "tm.event = 'Tx'"
	// EventChannelCapacity is a buffer size of channel between node and this program
	EventChannelCapacity = 2000

	PendingRequestsQueryPath = "/oracle.v1.Query/PendingRequests"
)

const (
	DefaultChainID          = "odin"
	DefaultNodeURL          = "tcp://localhost:26657"
	DefaultValidator        = ""
	DefaultExecutor         = ""
	DefaultGasPrices        = "10loki"
	DefaultLogLevel         = "info"
	DefaultBroadcastTimeout = "5m"
	DefaultRPCPollInterval  = "1s"
	DefaultMaxTry           = 5
	DefaultMaxReport        = 10
)

func runImpl(c *Context, l *Logger) error {
	l.Info(":rocket: Starting WebSocket subscriber")
	err := c.client.Start()
	if err != nil {
		return sdkerrors.Wrap(err, "failed to start websocket subscriber")
	}

	ctx, cxl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cxl()

	l.Info(":ear: Subscribing to events with query: %s...", TxQuery)
	eventChan, err := c.client.Subscribe(ctx, "", TxQuery, EventChannelCapacity)
	if err != nil {
		return err
	}

	if c.metricsEnabled {
		l.Info(":eyes: Starting Prometheus listener")
		go metricsListen(yoda.config.MetricsListenAddr, c)
	}

	availiableKeys := make([]bool, len(c.keys))
	waitingMsgs := make([][]ReportMsgWithKey, len(c.keys))
	for i := range availiableKeys {
		availiableKeys[i] = true
		waitingMsgs[i] = []ReportMsgWithKey{}
	}

	bz := cdc.MustMarshalBinaryBare(&oracletypes.QueryPendingRequestsRequest{
		ValidatorAddress: c.validator.String(),
	})
	resBz, err := c.client.ABCIQuery(context.Background(), PendingRequestsQueryPath, bz)
	if err != nil {
		l.Error(":exploding_head: Failed to get pending requests with error: %s", c, err.Error())
	}
	pendingRequests := oracletypes.QueryPendingRequestsResponse{}
	cdc.MustUnmarshalBinaryBare(resBz.Response.Value, &pendingRequests)

	l.Info(":mag: Found %d pending requests", len(pendingRequests.RequestIDs))
	for _, id := range pendingRequests.RequestIDs {
		c.pendingRequests[oracletypes.RequestID(id)] = true
		go handlePendingRequest(c, l.With("rid", id), oracletypes.RequestID(id))
	}

	for {
		select {
		case ev := <-eventChan:
			go handleTransaction(c, l, ev.Data.(tmtypes.EventDataTx).TxResult)
		case keyIndex := <-c.freeKeys:
			if len(waitingMsgs[keyIndex]) != 0 {
				if uint64(len(waitingMsgs[keyIndex])) > c.maxReport {
					go SubmitReport(c, l, keyIndex, waitingMsgs[keyIndex][:c.maxReport])
					waitingMsgs[keyIndex] = waitingMsgs[keyIndex][c.maxReport:]
				} else {
					go SubmitReport(c, l, keyIndex, waitingMsgs[keyIndex])
					waitingMsgs[keyIndex] = []ReportMsgWithKey{}
				}
			} else {
				availiableKeys[keyIndex] = true
			}
		case pm := <-c.pendingMsgs:
			c.updatePendingGauge(1)
			if availiableKeys[pm.keyIndex] {
				availiableKeys[pm.keyIndex] = false
				go SubmitReport(c, l, pm.keyIndex, []ReportMsgWithKey{pm})
			} else {
				waitingMsgs[pm.keyIndex] = append(waitingMsgs[pm.keyIndex], pm)
			}
		}
	}
}

func runCmd(ctx *Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{"r"},
		Short:   "Run the oracle process",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if yoda.config.ChainID == "" {
				return sdkerrors.Wrap(errors.ErrEmptyChainIDParam, "failed to parse chain in")
			}
			keys, err := yoda.keybase.List()
			if err != nil {
				return sdkerrors.Wrap(err, "failed to get keys from keyring")
			}
			if len(keys) == 0 {
				return sdkerrors.Wrap(errors.ErrNoKeyAvailable, "failed to get keys from keyring")
			}
			ctx.validator, err = sdk.ValAddressFromBech32(yoda.config.Validator)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to parse validator address")
			}
			err = sdk.VerifyAddressFormat(ctx.validator)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to verify address format")
			}
			allowLevel, err := log.AllowLevel(yoda.config.LogLevel)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to get log level")
			}
			logger := NewLogger(allowLevel)
			ctx.executor, err = executor.NewExecutor(yoda.config.Executor)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to create a new executor")
			}
			logger.Info(":star: Creating HTTP client with node URI: %s", yoda.config.NodeURI)

			ctx.client, err = httpclient.New(yoda.config.NodeURI, "/websocket")
			if err != nil {
				return sdkerrors.Wrap(err, "failed to create rpc client")
			}
			ctx.fileCache = filecache.New(filepath.Join(viper.GetString(flags.FlagHome), "files"))
			ctx.broadcastTimeout, err = time.ParseDuration(yoda.config.BroadcastTimeout)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to parse broadcast timeout")
			}
			ctx.keys = keys
			ctx.gasPrices = yoda.config.GasPrices
			ctx.maxTry = yoda.config.MaxTry
			ctx.maxReport = yoda.config.MaxReport
			ctx.rpcPollInterval, err = time.ParseDuration(yoda.config.RPCPollInterval)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to parse rpc poll interval")
			}
			ctx.pendingMsgs = make(chan ReportMsgWithKey)
			ctx.freeKeys = make(chan int64, len(keys))
			ctx.keyRoundRobinIndex = -1
			ctx.dataSourceCache = new(sync.Map)
			ctx.pendingRequests = make(map[oracletypes.RequestID]bool)
			ctx.metricsEnabled = yoda.config.MetricsListenAddr != ""
			return runImpl(ctx, logger)
		},
	}

	cmd.Flags().String(flags.FlagChainID, DefaultChainID, "chain ID of Odin network")
	if err := viper.BindPFlag(flags.FlagChainID, cmd.Flags().Lookup(flags.FlagChainID)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to parse %s flag", flags.FlagChainID))
	}
	cmd.Flags().String(flags.FlagNode, DefaultNodeURL, "RPC url to Odin node")
	if err := viper.BindPFlag(flags.FlagNode, cmd.Flags().Lookup(flags.FlagNode)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flags.FlagNode))
	}
	cmd.Flags().String(flagValidator, DefaultValidator, "validator address")
	if err := viper.BindPFlag(flagValidator, cmd.Flags().Lookup(flagValidator)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagValidator))
	}
	cmd.Flags().String(flagExecutor, DefaultExecutor, "executor name and url for executing the data source script")
	if err := viper.BindPFlag(flagExecutor, cmd.Flags().Lookup(flagExecutor)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagExecutor))
	}
	cmd.Flags().String(flags.FlagGasPrices, DefaultGasPrices, "gas prices for report transaction")
	if err := viper.BindPFlag(flags.FlagGasPrices, cmd.Flags().Lookup(flags.FlagGasPrices)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flags.FlagGasPrices))
	}
	cmd.Flags().String(flagLogLevel, DefaultLogLevel, "set the logger level")
	if err := viper.BindPFlag(flagLogLevel, cmd.Flags().Lookup(flagLogLevel)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagLogLevel))
	}
	cmd.Flags().String(flagBroadcastTimeout, DefaultBroadcastTimeout, "The time that Yoda will wait for tx commit")
	if err := viper.BindPFlag(flagBroadcastTimeout, cmd.Flags().Lookup(flagBroadcastTimeout)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagBroadcastTimeout))
	}
	cmd.Flags().String(flagRPCPollInterval, DefaultRPCPollInterval, "The duration of rpc poll interval")
	if err := viper.BindPFlag(flagRPCPollInterval, cmd.Flags().Lookup(flagRPCPollInterval)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagRPCPollInterval))
	}
	cmd.Flags().Uint64(flagMaxTry, DefaultMaxTry, "The maximum number of tries to submit a report transaction")
	if err := viper.BindPFlag(flagMaxTry, cmd.Flags().Lookup(flagMaxTry)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagMaxTry))
	}
	cmd.Flags().Uint64(flagMaxReport, DefaultMaxReport, "The maximum number of reports in one transaction")
	if err := viper.BindPFlag(flagMaxReport, cmd.Flags().Lookup(flagMaxReport)); err != nil {
		panic(sdkerrors.Wrapf(err, "failed to bind %s flag", flagMaxReport))
	}

	return cmd
}
