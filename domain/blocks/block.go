package blocks

import (
	"math/big"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type block struct {
	hash  hash.Hash
	body  Body
	proof *big.Int
}

func createBlock(
	hash hash.Hash,
	body Body,
	proof *big.Int,
) Block {
	out := block{
		hash:  hash,
		body:  body,
		proof: proof,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Body returns the body
func (obj *block) Body() Body {
	return obj.body
}

// Proof returns the proof
func (obj *block) Proof() *big.Int {
	return obj.proof
}
