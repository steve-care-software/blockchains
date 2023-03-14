package transactions

import (
	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// EnterOnCreateTransactionFn represents the enter's onCreate trx func
type EnterOnCreateTransactionFn func(body transactions.Body) (transactions.Transaction, error)

// ExitOnCreateTransactionFn represents the exit's onCreate trx func
type ExitOnCreateTransactionFn func(trx transactions.Transaction) error

// Builder represents the transaction's application builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	Now() (Application, error)
}

// Application represents a transaction application
type Application interface {
	List() ([]hash.Hash, error)
	Insert(trx transactions.Body) error
	InsertList(list []transactions.Body) error
	Retrieve(hash hash.Hash) (transactions.Transaction, error)
}
