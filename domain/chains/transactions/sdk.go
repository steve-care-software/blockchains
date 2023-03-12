package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewTransactionBuilder creates a new transaction builder
func NewTransactionBuilder() TransactionBuilder {
	hashAdapter := hash.NewAdapter()
	return createTransactionBuilder(hashAdapter)
}

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

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a transaction repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Transaction, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a transaction service
type Service interface {
	Insert(trx Transaction) error
	InsertList(list []Transaction) error
}
