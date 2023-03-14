package blocks

import (
	"math/big"
	"time"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewBodyBuilder creates a new body builder instance
func NewBodyBuilder() BodyBuilder {
	hashAdapter := hash.NewAdapter()
	return createBodyBuilder(hashAdapter)
}

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithBody(body Body) Builder
	WithProof(proof *big.Int) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Body() Body
	Proof() *big.Int
}

// BodyBuilder represents the body builder
type BodyBuilder interface {
	Create() BodyBuilder
	WithAddress(address []byte) BodyBuilder
	WithTransactions(trx transactions.Transactions) BodyBuilder
	WithParent(parent hash.Hash) BodyBuilder
	CreatedOn(createdOn time.Time) BodyBuilder
	Now() (Body, error)
}

// Body represents the block body
type Body interface {
	Hash() hash.Hash
	Address() []byte
	Transactions() transactions.Transactions
	CreatedOn() time.Time
	HasParent() bool
	Parent() *hash.Hash
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	WithKind(kind uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a block repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Block, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a block service
type Service interface {
	Insert(block Block) error
}
