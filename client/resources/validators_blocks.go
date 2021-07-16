package resources

type ValidatorBlocks struct {
	ValidatorAddress string `json:"validator_address"`
	BlocksCount      uint64 `json:"blocks_count"`
	StakePercentage  string `json:"stake_percentage"`
}

type BlockValidatorsResponse struct {
	ValidatorsBlocks []ValidatorBlocks `json:"validators_blocks"`
}
