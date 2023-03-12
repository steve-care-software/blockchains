package transactions

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type bodyBuilder struct {
	hashAdapter hash.Adapter
	address     []byte
	pFees       *uint
	scripts     []hash.Hash
}

func createBodyBuilder(
	hashAdapter hash.Adapter,
) BodyBuilder {
	out := bodyBuilder{
		hashAdapter: hashAdapter,
		address:     nil,
		pFees:       nil,
		scripts:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *bodyBuilder) Create() BodyBuilder {
	return createBodyBuilder(app.hashAdapter)
}

// WithAddress adds an address to the builder
func (app *bodyBuilder) WithAddress(address []byte) BodyBuilder {
	app.address = address
	return app
}

// WithFees add fees to the builder
func (app *bodyBuilder) WithFees(fees uint) BodyBuilder {
	app.pFees = &fees
	return app
}

// WithScripts add scripts to the builder
func (app *bodyBuilder) WithScripts(scripts []hash.Hash) BodyBuilder {
	app.scripts = scripts
	return app
}

// Now builds a new Body instance
func (app *bodyBuilder) Now() (Body, error) {
	if app.address != nil && len(app.address) <= 0 {
		app.address = nil
	}

	if app.address == nil {
		return nil, errors.New("the address is mandatory in order to build a transaction's Body instance")
	}

	if app.scripts != nil && len(app.scripts) <= 0 {
		app.scripts = nil
	}

	if app.scripts == nil {
		return nil, errors.New("the scripts are mandatory in order to build a transaction's Body instance")
	}

	if app.pFees == nil {
		return nil, errors.New("the fees is mandatory in order to build a transaction's Body instance")
	}

	data := [][]byte{
		app.address,
		[]byte(fmt.Sprintf("%d", *app.pFees)),
	}

	for _, oneScript := range app.scripts {
		data = append(data, oneScript.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createBody(*pHash, app.address, *app.pFees, app.scripts), nil
}
