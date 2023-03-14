package blocks

import (
	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// EnterOnBlockFn represents the enter's onBlock func
type EnterOnBlockFn func(block blocks.Block) error

// ExitOnBlockFn represents the exit's onBlock func
type ExitOnBlockFn func(block blocks.Block) error

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
