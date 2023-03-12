package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

type transaction struct {
	hash      hash.Hash
	body      Body
	signature []byte
}

func createTransaction(
	hash hash.Hash,
	body Body,
	signature []byte,
) Transaction {
	out := transaction{
		hash:      hash,
		body:      body,
		signature: signature,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Body returns the body
func (obj *transaction) Body() Body {
	return obj.body
}

// Signature returns the signature
func (obj *transaction) Signature() []byte {
	return obj.signature
}
