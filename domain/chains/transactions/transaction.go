package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

type transaction struct {
	hash      hash.Hash
	body      Body
	signature []byte
	pubKey    []byte
}

func createTransaction(
	hash hash.Hash,
	body Body,
	signature []byte,
	pubKey []byte,
) Transaction {
	out := transaction{
		hash:      hash,
		body:      body,
		signature: signature,
		pubKey:    pubKey,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// PublicKey returns the publicKey
func (obj *transaction) PublicKey() []byte {
	return obj.pubKey
}

// Body returns the body
func (obj *transaction) Body() Body {
	return obj.body
}

// Signature returns the signature
func (obj *transaction) Signature() []byte {
	return obj.signature
}
