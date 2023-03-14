package domain

import (
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithRoot(root hash.Hash) Builder
	WithHead(head blocks.Block) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	Name() string
	Root() hash.Hash
	HasHead() bool
	Head() blocks.Block
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	WithBlockKind(blockKind uint) RepositoryBuilder
	WithKind(kind uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a chain repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string) (Chain, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a chain service
type Service interface {
	Insert(chain Chain) error
}
