package handlers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

func GetChainIDHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.PostProcessResponseBare(w, clientCtx, map[string]string{"chain_id": clientCtx.ChainID})
	}
}

