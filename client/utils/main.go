package utils

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	tendermint "github.com/tendermint/tendermint/types"
	"time"
)

const (
	MaxCountPerPage = 100

	AllBlocksQuery = "block.height > 1"
	OrderByAsc     = "asc"
)

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
			return nil, sdkerrors.Wrap(err, "failed to find the blocks")
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

func TimeToUTCDate(t time.Time) time.Time {
	year, month, day := t.UTC().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func CalculateDays(t1, t2 time.Time) int {
	return int(t2.UTC().Sub(t1.UTC()).Hours())/24 + 1
}
