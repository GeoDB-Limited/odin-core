package rest

import (
	"fmt"
	"github.com/GeoDB-Limited/odin-core/client/rest/handlers"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

const (
	AppName   = "odin"
	RouterKey = AppName

	QueryChainID         = "chain_id"
	QueryGenesis         = "genesis"
	QueryAvgBlockSize    = "avg_block_size"
	QueryAvgBlockTime    = "avg_block_time"
	QueryDailyTxVolume   = "daily_tx_volume"
	QueryAvgTxFee        = "avg_tx_fee"
	QueryBlockValidators = "validators_blocks"
)

func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryChainID),
		handlers.GetChainIDHandler(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryGenesis),
		handlers.GetGenesisHandler(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryAvgBlockSize),
		handlers.GetAvgBlockSizeHandler(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryAvgBlockTime),
		handlers.GetAvgBlockTimeHandler(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryDailyTxVolume),
		handlers.GetDailyTxVolumeHandler(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryAvgTxFee),
		handlers.GetAvgTxFeeHandler(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/%s", RouterKey, QueryBlockValidators),
		handlers.GetValidatorsBlocksHandler(clientCtx),
	).Methods("GET")
}
