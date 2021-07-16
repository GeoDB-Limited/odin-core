package requests

import (
	"encoding/json"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"net/http"
	"time"
)

type ValidatorsBlocksRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Count     int64     `json:"count"`
}

func NewValidatorsBlocksRequest(r *http.Request) (*ValidatorsBlocksRequest, error) {
	req := ValidatorsBlocksRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}

func (r ValidatorsBlocksRequest) Validate() error {
	ok := r.StartDate.Before(r.EndDate)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid dates order")
	}
	if r.Count <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid validators count")
	}
	return nil
}
