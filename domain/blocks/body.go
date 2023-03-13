package blocks

import (
	"time"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type body struct {
	hash      hash.Hash
	address   []byte
	trx       transactions.Transactions
	createdOn time.Time
	pParent   *hash.Hash
}

func createBody(
	hash hash.Hash,
	address []byte,
	trx transactions.Transactions,
	createdOn time.Time,
) Body {
	return createBodyInternally(hash, address, trx, createdOn, nil)
}

func createBodyWithParent(
	hash hash.Hash,
	address []byte,
	trx transactions.Transactions,
	createdOn time.Time,
	pParent *hash.Hash,
) Body {
	return createBodyInternally(hash, address, trx, createdOn, pParent)
}

func createBodyInternally(
	hash hash.Hash,
	address []byte,
	trx transactions.Transactions,
	createdOn time.Time,
	pParent *hash.Hash,
) Body {
	out := body{
		hash:      hash,
		address:   address,
		trx:       trx,
		createdOn: createdOn,
		pParent:   pParent,
	}

	return &out
}

// Hash returns the hash
func (obj *body) Hash() hash.Hash {
	return obj.hash
}

// Address returns the address
func (obj *body) Address() []byte {
	return obj.address
}

// Transactions returns the transactions
func (obj *body) Transactions() transactions.Transactions {
	return obj.trx
}

// CreatedOn returns the creation time
func (obj *body) CreatedOn() time.Time {
	return obj.createdOn
}

// HasParent retruns true if there is parent, false otherwise
func (obj *body) HasParent() bool {
	return obj.pParent != nil
}

// Parent returns the parent, if any
func (obj *body) Parent() *hash.Hash {
	return obj.pParent
}
