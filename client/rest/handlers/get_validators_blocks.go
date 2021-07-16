package handlers

import (
	"context"
	"github.com/GeoDB-Limited/odin-core/client/resources"
	"github.com/GeoDB-Limited/odin-core/client/rest/requests"
	"github.com/GeoDB-Limited/odin-core/client/utils"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"net/http"
	"sort"
)

func GetValidatorsBlocksHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := requests.NewValidatorsBlocksRequest(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		node, err := clientCtx.GetNode()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctx := context.Background()
		validatorsBlocks := make(map[string]uint64, request.Count)

		blocksByDates, err := utils.GetBlocksByDates(clientCtx, request.StartDate, request.EndDate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if len(blocksByDates) == 0 {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Can not find the blocks")
			return
		}

		for _, blocks := range blocksByDates {
			for _, block := range blocks {
				validators, err := utils.GetBlockValidators(ctx, node, block.Header.Height)
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				for _, val := range validators {
					validatorsBlocks[val.Address.String()]++
				}
			}
		}

		stakingClient := stakingtypes.NewQueryClient(clientCtx)
		validatorBlocks := make([]resources.ValidatorBlocks, 0, len(validatorsBlocks))

		stakingPool, err := stakingClient.Pool(ctx, &stakingtypes.QueryPoolRequest{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		for k, v := range validatorsBlocks {
			vslAddr, err := sdk.ValAddressFromHex(k)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			req := &stakingtypes.QueryValidatorRequest{
				ValidatorAddr: vslAddr.String(),
			}
			res, err := stakingClient.Validator(ctx, req)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			stakePercentage := sdk.NewDecFromIntWithPrec(res.Validator.Tokens, 2).
				QuoRoundUp(sdk.NewDecFromIntWithPrec(stakingPool.Pool.BondedTokens, 2).
					Mul(sdk.NewDecFromIntWithPrec(sdk.NewInt(100), 2)))

			validatorBlocks = append(
				validatorBlocks,
				resources.ValidatorBlocks{
					ValidatorAddress: k,
					BlocksCount:      v,
					StakePercentage:  stakePercentage.String(),
				},
			)
		}

		sort.Slice(validatorBlocks, func(i, j int) bool {
			return validatorBlocks[i].BlocksCount > validatorBlocks[j].BlocksCount
		})

		validatorsCount := int64(len(validatorBlocks))
		if validatorsCount > request.Count {
			validatorsCount = request.Count
		}

		rest.PostProcessResponseBare(
			w,
			clientCtx,
			resources.BlockValidatorsResponse{ValidatorsBlocks: validatorBlocks[:validatorsCount]},
		)
	}
}
