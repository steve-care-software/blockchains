package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBodyBuilder creates a new body builder instance
func NewBodyBuilder() BodyBuilder {
	hashAdapter := hash.NewAdapter()
	return createBodyBuilder(hashAdapter)
}

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithBody(body Body) TransactionBuilder
	WithSignature(signature []byte) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Body() Body
	Signature() []byte
}

// BodyBuilder represents a body builder
type BodyBuilder interface {
	Create() BodyBuilder
	WithAddress(address []byte) BodyBuilder
	WithFees(fees uint) BodyBuilder
	WithScripts(scripts []hash.Hash) BodyBuilder
	Now() (Body, error)
}

// Body represents the transaction body
type Body interface {
	Hash() hash.Hash
	Address() []byte
	Fees() uint
	Scripts() []hash.Hash
}
