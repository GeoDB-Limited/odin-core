package handlers

import (
	"github.com/GeoDB-Limited/odin-core/client/resources"
	"github.com/GeoDB-Limited/odin-core/client/rest/requests"
	"github.com/GeoDB-Limited/odin-core/client/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

func GetAvgBlockTimeHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := requests.NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		averageBlockTime := make([]resources.AverageBlockTime, 0, utils.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := utils.GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
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
			averageBlockTime = append(averageBlockTime, resources.AverageBlockTime{
				Date:      key,
				BlockTime: uint64(averageTime),
			})
		}

		if len(averageBlockTime) == 0 {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Failed to find stats sata")
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, resources.AvgBlockTimeResponse{AverageBlockTime: averageBlockTime})
	}
}
