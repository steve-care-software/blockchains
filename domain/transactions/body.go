package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

type body struct {
	hash      hash.Hash
	address   hash.Hash
	fees      uint
	reference hash.Hash
}

func createBody(
	hash hash.Hash,
	address hash.Hash,
	fees uint,
	reference hash.Hash,
) Body {
	out := body{
		hash:      hash,
		address:   address,
		fees:      fees,
		reference: reference,
	}

	return &out
}

// Hash returns the hash
func (obj *body) Hash() hash.Hash {
	return obj.hash
}

// Address returns the address
func (obj *body) Address() hash.Hash {
	return obj.address
}

// Fees returns the fees
func (obj *body) Fees() uint {
	return obj.fees
}

// Reference returns the reference
func (obj *body) Reference() hash.Hash {
	return obj.reference
}
