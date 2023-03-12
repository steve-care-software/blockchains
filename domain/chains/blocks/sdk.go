package blocks

import (
	"math/big"
	"time"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

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
	Result() hash.Hash
	Difficulty() uint
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
	Parent() hash.Hash
}
