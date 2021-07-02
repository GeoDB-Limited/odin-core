package rpc

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidDateOrder = sdkerrors.Register(AppName, 1, "the end date is earlier than the start date")
)
