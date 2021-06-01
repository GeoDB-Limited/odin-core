package keeper

import (
	"fmt"
	auctiontypes "github.com/GeoDB-Limited/odin-core/x/auction/types"
	coinswaptypes "github.com/GeoDB-Limited/odin-core/x/coinswap/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	storeKey       sdk.StoreKey
	cdc            codec.BinaryMarshaler
	paramstore     paramstypes.Subspace
	oracleKeeper   auctiontypes.OracleKeeper
	coinswapKeeper auctiontypes.CoinswapKeeper
}

func NewKeeper(
	cdc codec.BinaryMarshaler,
	key sdk.StoreKey,
	subspace paramstypes.Subspace,
	ok auctiontypes.OracleKeeper,
	ck auctiontypes.CoinswapKeeper,
) Keeper {
	if !subspace.HasKeyTable() {
		subspace = subspace.WithKeyTable(auctiontypes.ParamKeyTable())
	}

	return Keeper{
		cdc:            cdc,
		storeKey:       key,
		paramstore:     subspace,
		oracleKeeper:   ok,
		coinswapKeeper: ck,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", auctiontypes.ModuleName))
}

// SetParams saves the given key-value parameter to the store.
func (k Keeper) SetParams(ctx sdk.Context, value auctiontypes.Params) {
	k.paramstore.SetParamSet(ctx, &value)
}

// GetParams returns all current parameters as a types.Params instance.
func (k Keeper) GetParams(ctx sdk.Context) (params auctiontypes.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// GetAuctionStartThreshold returns auction threshold parameter
func (k Keeper) GetAuctionStartThreshold(ctx sdk.Context) (res sdk.Coins) {
	k.paramstore.Get(ctx, auctiontypes.KeyAuctionStartThreshold, &res)
	return res
}

// GetExchangeRate returns auction exchange parameter
func (k Keeper) GetExchangeRate(ctx sdk.Context) (res coinswaptypes.Exchange) {
	k.paramstore.Get(ctx, auctiontypes.KeyExchangeRate, &res)
	return res
}

func (k Keeper) SetAuctionStatus(ctx sdk.Context, status auctiontypes.AuctionStatus) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&status)
	store.Set(auctiontypes.AuctionStatusStoreKey, b)
}

func (k Keeper) GetAuctionStatus(ctx sdk.Context) (payments auctiontypes.AuctionStatus) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(auctiontypes.AuctionStatusStoreKey)
	k.cdc.MustUnmarshalBinaryBare(bz, &payments)
	return
}

// GetAccumulatedPaymentsForData returns accumulated payments for data
func (k Keeper) GetAccumulatedPaymentsForData(ctx sdk.Context) sdk.Coins {
	accumulatedPaymentsForData := k.oracleKeeper.GetAccumulatedPaymentsForData(ctx)
	return accumulatedPaymentsForData.AccumulatedAmount
}

// StartAuction resolves to sell minigeo
func (k Keeper) StartAuction(ctx sdk.Context) error {
	status := k.GetAuctionStatus(ctx)
	if status.Pending {
		return sdkerrors.Wrap(auctiontypes.ErrAuctionHasStarted, "the auction is pending")
	}

	if err := k.coinswapKeeper.AddExchangeRate(ctx, k.GetExchangeRate(ctx)); err != nil {
		return sdkerrors.Wrap(err, "failed to start auction")
	}

	status.Pending = true
	k.SetAuctionStatus(ctx, status)

	return nil
}

// FinishAuction prohibits selling minigeo
func (k Keeper) FinishAuction(ctx sdk.Context) error {
	status := k.GetAuctionStatus(ctx)
	if !status.Pending {
		return sdkerrors.Wrap(auctiontypes.ErrAuctionHasClosed, "the auction is closed")
	}

	if err := k.coinswapKeeper.RemoveExchangeRate(ctx, k.GetExchangeRate(ctx)); err != nil {
		return sdkerrors.Wrap(err, "failed to remove auction exchange rate")
	}

	status.Pending = false
	k.SetAuctionStatus(ctx, status)

	return nil
}