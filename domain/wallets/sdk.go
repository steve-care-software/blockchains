package wallets

import (
	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Builder represents a wallet builder
type Builder interface {
	Create() Builder
	WithAddress(address []byte) Builder
	WithAmount(amount uint) Builder
	WithTransactions(trx transactions.Transactions) Builder
	Now() (Wallet, error)
}

// Wallet represents a wallet
type Wallet interface {
	Hash() hash.Hash
	Address() []byte
	Amount() uint
	Transactions() transactions.Transactions
}
