package transactions

import (
	"github.com/steve-care-software/blockchains/domain/chains"
	"github.com/steve-care-software/blockchains/domain/chains/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// EnterOnTrxFn represents the enter's onTrx func
type EnterOnTrxFn func(trx transactions.Transaction) error

// ExitOnTrxFn represents the exit's onTrx func
type ExitOnTrxFn func(trx transactions.Transaction) error

// Builder represents the transaction's application builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	Now() (Application, error)
}

// Application represents a transaction application
type Application interface {
	List() ([]hash.Hash, error)
	Insert(trx transactions.Transaction) error
	InsertList(list []transactions.Transaction) error
	Retrieve(hash hash.Hash) (transactions.Transaction, error)
}
