package transactions

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionBuilder struct {
	verifySigFn VerifySignatureFn
	hashAdapter hash.Adapter
	body        Body
	pubKey      []byte
	signature   []byte
}

func createTransactionBuilder(
	verifySigFn VerifySignatureFn,
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		verifySigFn: verifySigFn,
		hashAdapter: hashAdapter,
		body:        nil,
		pubKey:      nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder(
		app.verifySigFn,
		app.hashAdapter,
	)
}

// WithBody adds a body to the builder
func (app *transactionBuilder) WithBody(body Body) TransactionBuilder {
	app.body = body
	return app
}

// WithSignature adds a signature to the builder
func (app *transactionBuilder) WithSignature(signature []byte) TransactionBuilder {
	app.signature = signature
	return app
}

// WithPublicKey adds a pubKey to the builder
func (app *transactionBuilder) WithPublicKey(pubKey []byte) TransactionBuilder {
	app.pubKey = pubKey
	return app
}

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.body == nil {
		return nil, errors.New("the body is mandatory in order to build a Transaction instance")
	}

	if app.signature != nil && len(app.signature) <= 0 {
		app.signature = nil
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	if app.pubKey != nil && len(app.pubKey) <= 0 {
		app.pubKey = nil
	}

	if app.pubKey == nil {
		return nil, errors.New("the publicKey is mandatory in order to build a Transaction instance")
	}

	pPubKeyHash, err := app.hashAdapter.FromBytes(app.pubKey)
	if err != nil {
		return nil, err
	}

	if !app.body.Address().Compare(*pPubKeyHash) {
		str := fmt.Sprintf("the given publicKey, when hashed (%s) do not matched the body's Address (hash: %s) and therefore the Transaction is invalid", pPubKeyHash.String(), app.body.Address().String())
		return nil, errors.New(str)
	}

	if !app.verifySigFn(app.pubKey, app.signature, app.body.Hash()) {
		str := fmt.Sprintf("the signature could not be verified and therefore the Transaction is invalid")
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.body.Hash().Bytes(),
		app.signature,
		app.pubKey,
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*pHash, app.body, app.signature, app.pubKey), nil
}
