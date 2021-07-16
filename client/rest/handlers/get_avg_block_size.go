package handlers

import (
	"github.com/GeoDB-Limited/odin-core/client/resources"
	"github.com/GeoDB-Limited/odin-core/client/rest/requests"
	"github.com/GeoDB-Limited/odin-core/client/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

func GetAvgBlockSizeHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := requests.NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		averageBlockSize := make([]resources.AverageBlockSize, 0, utils.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := utils.GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
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
			averageBlockSize = append(averageBlockSize, resources.AverageBlockSize{
				Date: key,
				Size: uint64(averageSize),
			})
		}

		if len(averageBlockSize) == 0 {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Can not find stats data")
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, resources.AverageBlockSizeResponse{AverageBlockSize: averageBlockSize})
	}
}
