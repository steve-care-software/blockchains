package transactions

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionBuilder struct {
	hashAdapter hash.Adapter
	body        Body
	signature   []byte
}

func createTransactionBuilder(
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		hashAdapter: hashAdapter,
		body:        nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder(app.hashAdapter)
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

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.body == nil {
		return nil, errors.New("the body is mandatory in order to build a Transaction instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.body.Hash().Bytes(),
		app.signature,
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*pHash, app.body, app.signature), nil
}
