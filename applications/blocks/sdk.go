package blocks

import (
	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// EnterOnCreateBlockFn represents the enter's onCreate block func
type EnterOnCreateBlockFn func(abody blocks.Body) (blocks.Block, error)

// ExitOnCreateBlockFn represents the exit's onCreate block func
type ExitOnCreateBlockFn func(block blocks.Block) error

// Builder represents the block application builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	Now() (Application, error)
}

// Application represents the block application
type Application interface {
	List() ([]hash.Hash, error)
	Insert(block blocks.Block) error
	Retrieve(hash hash.Hash) (blocks.Block, error)
}
