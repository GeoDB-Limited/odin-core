package rest

import (
	"context"
	rpc "github.com/GeoDB-Limited/odin-core/client"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	tendermint "github.com/tendermint/tendermint/types"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const (
	MaxCountPerPage = 100

	AllBlocksQuery = "block.height > 1"
	OrderByAsc     = "asc"

	ValidatorsCountTag = "validators_count"
)

func GetChainIDFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.PostProcessResponseBare(w, clientCtx, map[string]string{"chain_id": clientCtx.ChainID})
	}
}

func GetGenesisHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		node, err := clientCtx.GetNode()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		genesis, err := node.Genesis(context.Background())
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponseBare(w, clientCtx, genesis.Genesis)
	}
}

type AverageBlockSize struct {
	Date time.Time `json:"date"`
	Size uint64    `json:"size"` // in bytes
}

type AverageBlockSizeResponse struct {
	AverageBlockSize []AverageBlockSize `json:"avg_block_size"`
}

func GetAvgBlockSize(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		averageBlockSize := make([]AverageBlockSize, 0, rpc.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		for key, value := range blocksByDates {
			totalSize := 0
			for _, block := range value {
				totalSize += block.Size()
			}
			averageSize := totalSize / len(value)
			averageBlockSize = append(averageBlockSize, AverageBlockSize{
				Date: key,
				Size: uint64(averageSize),
			})
		}

		rest.PostProcessResponseBare(w, clientCtx, AverageBlockSizeResponse{averageBlockSize})
	}
}

type AverageBlockTime struct {
	Date      time.Time `json:"date"`
	BlockTime uint64    `json:"block_time"` // in seconds
}

type AvgBlockTimeResponse struct {
	AverageBlockTime []AverageBlockTime `json:"avg_block_time"`
}

func GetAvgBlockTime(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		averageBlockTime := make([]AverageBlockTime, 0, rpc.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		for key, value := range blocksByDates {
			commonSeconds := 0

			blocksCount := len(value)
			for i := 1; i < blocksCount; i++ {
				commonSeconds += int(value[i].Header.Time.Unix() - value[i-1].Header.Time.Unix())
			}

			averageTime := commonSeconds / len(value)
			averageBlockTime = append(averageBlockTime, AverageBlockTime{
				Date:      key,
				BlockTime: uint64(averageTime),
			})
		}

		rest.PostProcessResponseBare(w, clientCtx, AvgBlockTimeResponse{averageBlockTime})
	}
}

type DailyTxVolume struct {
	Date   time.Time `json:"date"`
	Volume uint64    `json:"volume"`
}

type DailyTxVolumesResponse struct {
	DailyTxVolumes []DailyTxVolume `json:"daily_tx_volumes"`
}

func GetDailyTxVolume(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		dailyTxsVolumes := make([]DailyTxVolume, 0, rpc.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		for key, value := range blocksByDates {
			txsVolume := 0
			for _, block := range value {
				txsVolume += len(block.Data.Txs)
			}

			dailyTxsVolumes = append(dailyTxsVolumes, DailyTxVolume{
				Date:   key,
				Volume: uint64(txsVolume),
			})
		}

		rest.PostProcessResponseBare(w, clientCtx, DailyTxVolumesResponse{dailyTxsVolumes})
	}
}

type AverageTxFee struct {
	Date time.Time `json:"date"`
	Fee  sdk.Coins `json:"fee"`
}

type AverageTxFeesResponse struct {
	AverageTxFees []AverageTxFee `json:"avg_tx_fees"`
}

func GetAvgTxFee(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		averageTxFee := make([]AverageTxFee, 0, rpc.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		for key, value := range blocksByDates {
			totalFee := sdk.NewCoins()
			totalTxCount := 0
			for _, block := range value {
				for _, tx := range block.Data.Txs {
					tx, err := clientCtx.TxConfig.TxDecoder()(tx)
					if err != nil {
						rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
						return
					}

					feeTx, ok := tx.(sdk.FeeTx)
					if !ok {
						rest.WriteErrorResponse(w, http.StatusInternalServerError, sdkerrors.ErrTxDecode.Error())
						return
					}

					txFee := feeTx.GetFee()
					totalFee = totalFee.Add(txFee...)
				}
				totalTxCount += len(block.Data.Txs)
			}

			avgFee := sdk.NewCoins()
			if totalTxCount != 0 {
				avgFee, _ = sdk.NewDecCoinsFromCoins(totalFee...).QuoDec(sdk.NewDec(int64(totalTxCount))).TruncateDecimal()
			}
			averageTxFee = append(averageTxFee, AverageTxFee{
				Date: key,
				Fee:  avgFee,
			})
		}

		rest.PostProcessResponseBare(w, clientCtx, AverageTxFeesResponse{averageTxFee})
	}
}

type ValidatorBlocks struct {
	ValidatorAddress string `json:"validator_address"`
	BlocksCount      uint64 `json:"blocks_count"`
}

type BlockValidatorsResponse struct {
	ValidatorsBlocks []ValidatorBlocks `json:"validators_blocks"`
}

func GetBlockValidatorsStat(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validatorsCount, err := strconv.ParseInt(mux.Vars(r)[ValidatorsCountTag], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse block height")
			return
		}

		node, err := clientCtx.GetNode()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		blockValidators := make(map[string]uint64, validatorsCount)
		ctx := context.Background()

		info, err := node.ABCIInfo(context.Background())
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		height := int(info.Response.LastBlockHeight)
		for i := 1; i <= height; i++ {
			blockHeight := int64(i)
			validators, err := GetBlockValidators(ctx, node, blockHeight)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			for _, val := range validators {
				blockValidators[val.Address.String()]++
			}
		}

		validatorBlocks := make([]ValidatorBlocks, 0, len(blockValidators))
		for k, v := range blockValidators {
			validatorBlocks = append(validatorBlocks, ValidatorBlocks{k, v})
		}

		sort.Slice(validatorBlocks, func(i, j int) bool {
			return validatorBlocks[i].BlocksCount > validatorBlocks[j].BlocksCount
		})

		length := int64(len(validatorBlocks))
		if length < validatorsCount {
			validatorsCount = length
		}

		rest.PostProcessResponseBare(w, clientCtx, BlockValidatorsResponse{validatorBlocks[:validatorsCount]})
	}
}

func GetBlocksByDates(clientCtx client.Context, startDate, endDate time.Time) (map[time.Time][]*tendermint.Block, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get a node")
	}

	info, err := node.ABCIInfo(context.Background())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get ABCI info")
	}

	maxBlocksPerPage := MaxCountPerPage
	height := int(info.Response.LastBlockHeight)
	pages := height / maxBlocksPerPage
	if pages*maxBlocksPerPage < height {
		pages++
	}

	blocksPerDay := make(map[time.Time][]*tendermint.Block, 0)

	for i := 1; i <= pages; i++ {
		pageNumber := i
		result, err := node.BlockSearch(
			context.Background(),
			AllBlocksQuery,
			&pageNumber,
			&maxBlocksPerPage,
			OrderByAsc,
		)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "failed to get find the blocks")
		}

		for _, r := range result.Blocks {
			blockDate := rpc.TimeToUTCDate(r.Block.Header.Time)
			if blockDate.Equal(startDate) || blockDate.After(startDate) && blockDate.Equal(endDate) || blockDate.Before(endDate) {
				blocksPerDay[blockDate] = append(blocksPerDay[blockDate], r.Block)
			}
		}
	}

	return blocksPerDay, nil
}

func GetBlockValidators(ctx context.Context, node rpcclient.Client, blockHeight int64) ([]tendermint.Validator, error) {
	var validators []tendermint.Validator
	maxValidatorsPerPage := MaxCountPerPage
	page := 1

	for {
		resultValidators, err := node.Validators(ctx, &blockHeight, &page, &maxValidatorsPerPage)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "failed to get a node")
		}

		for _, val := range resultValidators.Validators {
			validators = append(validators, *val)
		}

		if resultValidators.Total == len(validators) {
			break
		}

		page++
	}

	return validators, nil
}
