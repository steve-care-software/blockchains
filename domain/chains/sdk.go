package chains

import (
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
)

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
	Head() blocks.Block
}
