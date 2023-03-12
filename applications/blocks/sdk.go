package blocks

import (
	"github.com/steve-care-software/blockchains/domain/chains"
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// ApplicationOnBlockFn represents the application's onBlock func
type ApplicationOnBlockFn func(block blocks.Block) error

// Builder represents the block application builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	Now() (Application, error)
}

// Application represents the block application
type Application interface {
	List() ([]hash.Hash, error)
	Mine(trx transactions.Transactions) error
	Retrieve(hash hash.Hash) (blocks.Block, error)
}
