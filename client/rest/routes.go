package rest

import (
	"fmt"
	"github.com/GeoDB-Limited/odin-core/client"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/chain_id", rpc.RouterKey), GetChainIDFn(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/genesis", rpc.RouterKey), GetGenesisHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/avg_block_size", rpc.RouterKey), GetAvgBlockSize(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/avg_block_time", rpc.RouterKey), GetAvgBlockTime(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/daily_tx_volume", rpc.RouterKey), GetDailyTxVolume(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/avg_tx_fee", rpc.RouterKey), GetAvgTxFee(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/block_validators/{%s}", rpc.RouterKey, ValidatorsCountTag), GetBlockValidatorsStat(clientCtx)).Methods("GET")
}
