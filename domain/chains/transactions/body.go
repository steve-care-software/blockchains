package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

type body struct {
	hash    hash.Hash
	address []byte
	fees    uint
	scripts []hash.Hash
}

func createBody(
	hash hash.Hash,
	address []byte,
	fees uint,
	scripts []hash.Hash,
) Body {
	out := body{
		hash:    hash,
		address: address,
		fees:    fees,
		scripts: scripts,
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

// Fees returns the fees
func (obj *body) Fees() uint {
	return obj.fees
}

// Scripts returns the scripts
func (obj *body) Scripts() []hash.Hash {
	return obj.scripts
}
