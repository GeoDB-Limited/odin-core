package resources

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type AverageTxFee struct {
	Date time.Time `json:"date"`
	Fee  sdk.Coins `json:"fee"`
}

type AverageTxFeesResponse struct {
	AverageTxFees []AverageTxFee `json:"avg_tx_fees"`
}

