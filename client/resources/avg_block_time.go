package resources

import "time"

type AverageBlockTime struct {
	Date      time.Time `json:"date"`
	BlockTime uint64    `json:"block_time"` // in seconds
}

type AvgBlockTimeResponse struct {
	AverageBlockTime []AverageBlockTime `json:"avg_block_time"`
}

