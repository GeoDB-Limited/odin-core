package resources

import "time"

type DailyTxVolume struct {
	Date   time.Time `json:"date"`
	Volume uint64    `json:"volume"`
}

type DailyTxVolumesResponse struct {
	DailyTxVolumes []DailyTxVolume `json:"daily_tx_volumes"`
}
