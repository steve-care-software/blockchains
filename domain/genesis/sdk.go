package genesis

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

const maxMiningValue = 9

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents the genesis builder
type Builder interface {
	Create() Builder
	WithDifficulty(difficulty uint) Builder
	WithReward(reward uint) Builder
	WithHalving(halving uint) Builder
	WithMiningValue(miningValue uint8) Builder
	WithMaxTrxPerBlock(maxTrxPerBlock uint) Builder
	WithBlockDuration(blockDuration time.Duration) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents the genesis block
type Genesis interface {
	Hash() hash.Hash
	Difficulty() uint
	Reward() uint
	Halving() uint
	MiningValue() uint8
	MaxTrxPerBlock() uint
	BlockDuration() time.Duration
	CreatedOn() time.Time
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a genesis repository
type Repository interface {
	Retrieve(hash hash.Hash) (Genesis, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a genesis service
type Service interface {
	Insert(genesis Genesis) error
}
