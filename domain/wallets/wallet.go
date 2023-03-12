package wallets

import (
	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type wallet struct {
	hash    hash.Hash
	address []byte
	amount  uint
	trx     transactions.Transactions
}

func createWallet(
	hash hash.Hash,
	address []byte,
	amount uint,
	trx transactions.Transactions,
) Wallet {
	out := wallet{
		hash:    hash,
		address: address,
		amount:  amount,
		trx:     trx,
	}

	return &out
}

// Hash returns the hash
func (obj *wallet) Hash() hash.Hash {
	return obj.hash
}

// Address returns the address
func (obj *wallet) Address() []byte {
	return obj.address
}

// Amount returns the amount
func (obj *wallet) Amount() uint {
	return obj.amount
}

// Transactions returns the transactions
func (obj *wallet) Transactions() transactions.Transactions {
	return obj.trx
}
