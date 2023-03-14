package domain

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	name string
	root hash.Hash
	head blocks.Block
}

func createBuilder() Builder {
	out := builder{
		name: "",
		root: nil,
		head: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root hash.Hash) Builder {
	app.root = root
	return app
}

// WithHead adds a head to the builder
func (app *builder) WithHead(head blocks.Block) Builder {
	app.head = head
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Chain instance")
	}

	if app.root == nil {
		return nil, errors.New("the genesis is mandatory in order to build a Chain instance")
	}

	if app.head != nil {
		return createChainWithHead(app.name, app.root, app.head), nil
	}

	return createChain(app.name, app.root), nil
}
