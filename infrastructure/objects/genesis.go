package objects

import "time"

// Genesis represents a genesis
type Genesis struct {
	Difficulty     uint          `json:"difficulty"`
	Reward         uint          `json:"reward"`
	Halving        uint          `json:"halving"`
	MiningValue    uint8         `json:"mining_value"`
	MaxTrxPerBlock uint          `json:"max_trx_per_block"`
	BlockDuration  time.Duration `json:"block_duration"`
	CreatedOn      time.Time     `json:"created_on"`
}
