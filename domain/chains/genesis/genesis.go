package genesis

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type genesis struct {
	hash           hash.Hash
	difficulty     uint
	reward         uint
	halving        uint
	miningValue    uint8
	maxTrxPerBlock uint
	blockDuration  time.Duration
	createdOn      time.Time
}

func createGenesis(
	hash hash.Hash,
	difficulty uint,
	reward uint,
	halving uint,
	miningValue uint8,
	maxTrxPerBlock uint,
	blockDuration time.Duration,
	createdOn time.Time,
) Genesis {
	out := genesis{
		hash:           hash,
		difficulty:     difficulty,
		reward:         reward,
		halving:        halving,
		miningValue:    miningValue,
		maxTrxPerBlock: maxTrxPerBlock,
		blockDuration:  blockDuration,
		createdOn:      createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Difficulty returns the difficulty
func (obj *genesis) Difficulty() uint {
	return obj.difficulty
}

// Reward returns the reward
func (obj *genesis) Reward() uint {
	return obj.reward
}

// Halving returns the halving
func (obj *genesis) Halving() uint {
	return obj.halving
}

// MiningValue returns the miningValue
func (obj *genesis) MiningValue() uint8 {
	return obj.miningValue
}

// MaxTrxPerBlock returns the maxTrxPerBlock
func (obj *genesis) MaxTrxPerBlock() uint {
	return obj.maxTrxPerBlock
}

// BlockDuration returns the blockDuration
func (obj *genesis) BlockDuration() time.Duration {
	return obj.blockDuration
}

// CreatedOn returns the creation time
func (obj *genesis) CreatedOn() time.Time {
	return obj.createdOn
}
