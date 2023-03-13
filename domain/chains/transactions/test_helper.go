package transactions

import "github.com/steve-care-software/libs/cryptography/hash"

// NewTransactionForTests creates a new transaction for tests
func NewTransactionForTests(pubKey []byte) Transaction {
	hashAdapter := hash.NewAdapter()
	pAddress, err := hashAdapter.FromBytes(pubKey)
	if err != nil {
		panic(err)
	}

	body := NewBodyForTests(*pAddress)
	signature := []byte("This is a signature")
	ins, err := NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithPublicKey(pubKey).WithSignature(signature).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBodyForTests creates a new body for tests
func NewBodyForTests(address hash.Hash) Body {
	hashAdapter := hash.NewAdapter()
	pReferenceHash, err := hashAdapter.FromBytes([]byte("this is a reference"))
	if err != nil {
		panic(err)
	}

	fees := uint(34)
	body, err := NewBodyBuilder().Create().WithAddress(address).WithReference(*pReferenceHash).WithFees(fees).Now()
	if err != nil {
		panic(err)
	}

	return body
}
