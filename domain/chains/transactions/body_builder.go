package transactions

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type bodyBuilder struct {
	hashAdapter hash.Adapter
	address     hash.Hash
	pFees       *uint
	reference   hash.Hash
}

func createBodyBuilder(
	hashAdapter hash.Adapter,
) BodyBuilder {
	out := bodyBuilder{
		hashAdapter: hashAdapter,
		address:     nil,
		pFees:       nil,
		reference:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *bodyBuilder) Create() BodyBuilder {
	return createBodyBuilder(app.hashAdapter)
}

// WithAddress adds an address to the builder
func (app *bodyBuilder) WithAddress(address hash.Hash) BodyBuilder {
	app.address = address
	return app
}

// WithFees add fees to the builder
func (app *bodyBuilder) WithFees(fees uint) BodyBuilder {
	app.pFees = &fees
	return app
}

// WithReference add reference to the builder
func (app *bodyBuilder) WithReference(reference hash.Hash) BodyBuilder {
	app.reference = reference
	return app
}

// Now builds a new Body instance
func (app *bodyBuilder) Now() (Body, error) {
	if app.address == nil {
		return nil, errors.New("the address is mandatory in order to build a transaction's Body instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference are mandatory in order to build a transaction's Body instance")
	}

	if app.pFees == nil {
		return nil, errors.New("the fees is mandatory in order to build a transaction's Body instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.address.Bytes(),
		app.reference.Bytes(),
		hash.Hash(fmt.Sprintf("%d", *app.pFees)),
	})

	if err != nil {
		return nil, err
	}

	return createBody(*pHash, app.address, *app.pFees, app.reference), nil
}
