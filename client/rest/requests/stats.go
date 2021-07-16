package requests

import (
	"encoding/json"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"net/http"
	"time"
)

type StatsRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func NewStatsRequest(r *http.Request) (*StatsRequest, error) {
	req := StatsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}

func (r StatsRequest) Validate() error {
	ok := r.StartDate.Before(r.EndDate)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid dates order")
	}
	return nil
}