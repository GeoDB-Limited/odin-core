package handlers

import (
	"github.com/GeoDB-Limited/odin-core/client/resources"
	"github.com/GeoDB-Limited/odin-core/client/rest/requests"
	"github.com/GeoDB-Limited/odin-core/client/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

func GetDailyTxVolumeHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := requests.NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		dailyTxsVolumes := make([]resources.DailyTxVolume, 0, utils.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := utils.GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if len(blocksByDates) == 0 {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Failed to find blocks")
			return
		}

		for key, value := range blocksByDates {
			txsVolume := 0
			for _, block := range value {
				txsVolume += len(block.Data.Txs)
			}

			dailyTxsVolumes = append(dailyTxsVolumes, resources.DailyTxVolume{
				Date:   key,
				Volume: uint64(txsVolume),
			})
		}

		if len(dailyTxsVolumes) == 0 {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Failed to find stats sata")
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, resources.DailyTxVolumesResponse{DailyTxVolumes: dailyTxsVolumes})
	}
}
