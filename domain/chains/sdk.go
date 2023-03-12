package chains

import (
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithRoot(root genesis.Genesis) Builder
	WithHead(head blocks.Block) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	Name() string
	Root() genesis.Genesis
	HasHead() bool
	Head() blocks.Block
}
