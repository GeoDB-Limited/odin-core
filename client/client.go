package rpc

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

const (
	RouterKey = "odin"
)

func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/chain_id", RouterKey), GetChainIDFn(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/genesis", RouterKey), GetGenesisHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/evm-validators", RouterKey), GetEVMValidators(clientCtx)).Methods("GET")
}
