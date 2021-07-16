package resources

import "time"

type AverageBlockSize struct {
	Date time.Time `json:"date"`
	Size uint64    `json:"size"` // in bytes
}

type AverageBlockSizeResponse struct {
	AverageBlockSize []AverageBlockSize `json:"avg_block_size"`
}
