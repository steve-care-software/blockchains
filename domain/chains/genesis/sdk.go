package genesis

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// Builder represents the genesis builder
type Builder interface {
	Create() Builder
	WithDifficulty(difficulty uint) Builder
	WithReward(reward uint) Builder
	WithHalving(halving uint) Builder
	WithMiningValue(miningValue uint8) Builder
	WithMinTrxPerBlock(minTrxPerBlock uint) Builder
	WithMaxTrxPerBlock(maxTrxPerBlock uint) Builder
	WithBlockDuration(blockDuration time.Duration) Builder
	WithInitTime(initTime time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents the genesis block
type Genesis interface {
	Hash() hash.Hash
	Difficulty() uint
	Reward() uint
	Halving() uint
	MiningValue() uint8
	MinTrxPerBlock() uint
	MaxTrxPerBlock() uint
	BlockDuration() time.Duration
	InitTime() time.Time
}
