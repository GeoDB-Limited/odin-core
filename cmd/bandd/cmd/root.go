package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	band "github.com/GeoDB-Limited/odin-core/app"
	"github.com/GeoDB-Limited/odin-core/app/params"
	"github.com/GeoDB-Limited/odin-core/hooks/emitter"
	"github.com/GeoDB-Limited/odin-core/hooks/request"
)

const (
	flagWithEmitter           = "with-emitter"
	flagDisableFeelessReports = "disable-feeless-reports"
	flagEnableFastSync        = "enable-fast-sync"
	flagWithPricer            = "with-pricer"
	flagWithRequestSearch     = "with-request-search"
	flagWithOwasmCacheSize    = "oracle-script-cache-size"
	flagEnableApi             = "api.enable"
)

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := band.MakeEncodingConfig()
	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(band.DefaultNodeHome)

	rootCmd := &cobra.Command{
		Use:   "bandd",
		Short: "Band Consumer App",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	authclient.Codec = encodingConfig.Marshaler

	rootCmd.AddCommand(
		InitCmd(band.NewDefaultGenesisState(), band.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, band.DefaultNodeHome),
		// band.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(band.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, band.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(band.ModuleBasics),
		AddGenesisAccountCmd(band.DefaultNodeHome),
		AddGenesisDataSourceCmd(band.DefaultNodeHome),
		AddGenesisOracleScriptCmd(band.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
	)

	server.AddCommands(rootCmd, band.DefaultNodeHome, newApp, createSimappAndExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(band.DefaultNodeHome),
	)

	rootCmd.PersistentFlags().String(flagWithRequestSearch, "", "[Experimental] Enable mode to save request in sql database")
	rootCmd.PersistentFlags().String(flagWithEmitter, "", "[Experimental] Enable mode with emitter")
	rootCmd.PersistentFlags().Uint32(flagWithOwasmCacheSize, 100, "[Experimental] Number of oracle scripts to cache")
}
func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	band.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
	)

	band.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// newApp is an AppCreator
func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	bandApp := band.NewBandApp(
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		band.MakeEncodingConfig(), // Ideally, we would reuse the one created by NewRootCmd.
		appOpts,
		cast.ToBool(appOpts.Get(flagDisableFeelessReports)),
		cast.ToUint32(appOpts.Get(flagWithOwasmCacheSize)),
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent))),
	)
	connStr, _ := appOpts.Get(flagWithRequestSearch).(string)
	if connStr != "" {
		bandApp.AddHook(request.NewHook(
			bandApp.AppCodec(), bandApp.OracleKeeper, connStr))
	}

	connStr, _ = appOpts.Get(flagWithEmitter).(string)
	if connStr != "" {
		bandApp.AddHook(
			emitter.NewHook(bandApp.AppCodec(), bandApp.LegacyAmino(), band.MakeEncodingConfig(), bandApp.AccountKeeper, bandApp.BankKeeper,
				bandApp.StakingKeeper, bandApp.MintKeeper, bandApp.DistrKeeper, bandApp.GovKeeper,
				bandApp.OracleKeeper, connStr, false))
	}

	return bandApp
}

func createSimappAndExport(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string,
	appOpts servertypes.AppOptions) (servertypes.ExportedApp, error) {

	encCfg := band.MakeEncodingConfig() // Ideally, we would reuse the one created by NewRootCmd.
	encCfg.Marshaler = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	var bandConsumerApp *band.BandApp
	if height != -1 {
		bandConsumerApp = band.NewBandApp(logger, db, traceStore, false, map[int64]bool{}, "", uint(1), encCfg, appOpts, false, cast.ToUint32(appOpts.Get(flagWithOwasmCacheSize)))

		if err := bandConsumerApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		bandConsumerApp = band.NewBandApp(logger, db, traceStore, true, map[int64]bool{}, "", uint(1), encCfg, appOpts, false, cast.ToUint32(appOpts.Get(flagWithOwasmCacheSize)))
	}

	return bandConsumerApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
