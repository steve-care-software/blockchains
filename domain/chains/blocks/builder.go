package blocks

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	body        Body
	proof       *big.Int
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		body:        nil,
		proof:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithBody adds a body to the builder
func (app *builder) WithBody(body Body) Builder {
	app.body = body
	return app
}

// WithProof adds a proof to the builder
func (app *builder) WithProof(proof *big.Int) Builder {
	app.proof = proof
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.body == nil {
		return nil, errors.New("the body is mandatory in order to build a Block instance")
	}

	if app.proof == nil {
		return nil, errors.New("the proof is mandatory in order to build a Block instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.body.Hash(),
		[]byte(app.proof.String()),
	})

	if err != nil {
		return nil, err
	}

	return createBlock(*pHash, app.body, app.proof), nil
}
