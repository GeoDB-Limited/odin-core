package handlers

import (
	"github.com/GeoDB-Limited/odin-core/client/resources"
	"github.com/GeoDB-Limited/odin-core/client/rest/requests"
	"github.com/GeoDB-Limited/odin-core/client/utils"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

func GetAvgTxFeeHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := requests.NewStatsRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		averageTxFee := make([]resources.AverageTxFee, 0, utils.CalculateDays(request.StartDate, request.EndDate))

		blocksByDates, err := utils.GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
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
			averageTxFee = append(averageTxFee, resources.AverageTxFee{
				Date: key,
				Fee:  avgFee,
			})
		}

		if len(averageTxFee) == 0 {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Failed to find stats sata")
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, resources.AverageTxFeesResponse{AverageTxFees: averageTxFee})
	}
}
