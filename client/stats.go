package rpc

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
	"time"
)

type BlockAverageBlockSizeRequest struct {
	StartDate time.Time `json:"start_time"`
	EndDate   time.Time `json:"end_time"`
}

type BlockAverageBlockSizeResponse struct {
	Result []AverageBlockSize `json:"average_block_size"`
}

type AverageBlockSize struct {
	Date        time.Time `json:"date"`
	AverageSize uint64    `json:"average_size"`
}

func GetAvgBlockSize(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		node, err := clientCtx.GetNode()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		req := BlockAverageBlockSizeRequest{
			StartDate: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Now().UTC(),
		}

		info, err := node.ABCIInfo(context.Background())
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		height := int(info.Response.LastBlockHeight)
		maxPerPage := 100
		pages := height / maxPerPage
		if pages*maxPerPage < height {
			pages += 1
		}

		query := "block.height > 0"

		var res []AverageBlockSize
		var blocksCount, commonSize uint64
		year, month, day := 0, time.Month(0), 0

		for i := 1; i <= pages; i++ {
			pageNumber := i
			result, err := node.BlockSearch(context.Background(), query, &pageNumber, &maxPerPage, "asc")
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			for j, r := range result.Blocks {
				blockTime := r.Block.Header.Time.UTC()
				if blockTime.After(req.StartDate) && blockTime.Before(req.EndDate) {
					y, m, d := blockTime.Date()

					isLastPage := j == len(result.Blocks)-1 && i == pages
					if y != year || m != month || d != day || isLastPage {
						if blocksCount != 0 || isLastPage {
							averageSize := commonSize / blocksCount
							date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
							bs := AverageBlockSize{
								AverageSize: averageSize,
								Date:        date,
							}
							res = append(res, bs)

							commonSize = 0
							blocksCount = 0
						}

						year, month, day = y, m, d
					}

					commonSize += uint64(r.Block.Size())
					blocksCount++
				}
			}
		}

		ro := BlockAverageBlockSizeResponse{
			Result: res,
		}

		rest.PostProcessResponseBare(w, clientCtx, ro)
	}
}
