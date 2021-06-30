package rpc

import (
	"context"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	tendermint "github.com/tendermint/tendermint/types"
	"net/http"
	"time"
)

const (
	MaxBlocksPerPage = 100

	AllBlocksQuery = "block.height > 0"
	OrderByAsc     = "asc"
)

type BlockAverageBlockSizeRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type AverageBlockSize struct {
	Date        time.Time `json:"date"`
	AverageSize uint64    `json:"bytes"`
}

type AverageTxsVolume struct {
	Date      time.Time `json:"date"`
	TxsVolume uint64    `json:"txs_volume"`
}

func GetAvgBlockSize(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request BlockAverageBlockSizeRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var averageBlockSize []AverageBlockSize

		blocksByDates, err := GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		for key, value := range blocksByDates {
			commonSize := 0
			for _, block := range value {
				commonSize += block.Size()
			}
			averageSize := commonSize / len(value)
			averageBlockSize = append(averageBlockSize, AverageBlockSize{
				Date:        key,
				AverageSize: uint64(averageSize),
			})
		}

		rest.PostProcessResponseBare(w, clientCtx, averageBlockSize)
	}
}

func GetDailyTxsVolume(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request BlockAverageBlockSizeRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var dailyTxsVolume []AverageTxsVolume

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

			dailyTxsVolume = append(dailyTxsVolume, AverageTxsVolume{
				Date:      key,
				TxsVolume: uint64(txsVolume),
			})
		}

		rest.PostProcessResponseBare(w, clientCtx, dailyTxsVolume)
	}
}

func GetBlocksByDates(clientCtx client.Context, startDate, endDate time.Time) (map[time.Time][]*tendermint.Block, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get node")
	}

	info, err := node.ABCIInfo(context.Background())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get ABCI info")
	}

	maxBlocksPerPage := MaxBlocksPerPage
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
			blockDate := TimeToUTCDate(r.Block.Header.Time)
			if blockDate.Equal(startDate) || blockDate.After(startDate) && blockDate.Equal(endDate) || blockDate.Before(endDate) {
				blocksPerDay[blockDate] = append(blocksPerDay[blockDate], r.Block)
			}
		}
	}

	return blocksPerDay, nil
}

func TimeToUTCDate(t time.Time) time.Time {
	year, month, day := t.UTC().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
