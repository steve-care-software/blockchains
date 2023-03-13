package blocks

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type bodyBuilder struct {
	hashAdapter hash.Adapter
	address     []byte
	trx         transactions.Transactions
	pCreatedOn  *time.Time
	pParent     *hash.Hash
}

func createBodyBuilder(
	hashAdapter hash.Adapter,
) BodyBuilder {
	out := bodyBuilder{
		hashAdapter: hashAdapter,
		address:     nil,
		trx:         nil,
		pCreatedOn:  nil,
		pParent:     nil,
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

// WithTransactions add transactions to the builder
func (app *bodyBuilder) WithTransactions(trx transactions.Transactions) BodyBuilder {
	app.trx = trx
	return app
}

// WithParent adds parent to the builder
func (app *bodyBuilder) WithParent(parent hash.Hash) BodyBuilder {
	app.pParent = &parent
	return app
}

// CreatedOn adds a creation time to the builder
func (app *bodyBuilder) CreatedOn(createdOn time.Time) BodyBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Body instance
func (app *bodyBuilder) Now() (Body, error) {
	if app.address != nil && len(app.address) <= 0 {
		app.address = nil
	}

	if app.address == nil {
		return nil, errors.New("the address is mandatory in order to build a block's Body instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions are mandatory in order to build a block's Body instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creationTime's timestamp is mandatory in order to build a block's Body instance")
	}

	if app.pCreatedOn.Unix() <= 0 {
		return nil, errors.New("the creationTime's timestamp must be greater than zero (0) in order to build a block's Body instance")
	}

	data := [][]byte{
		app.address,
		app.trx.Hash().Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	}

	if app.pParent != nil {
		data = append(data, app.pParent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pParent != nil {
		return createBodyWithParent(*pHash, app.address, app.trx, *app.pCreatedOn, app.pParent), nil
	}

	return createBody(*pHash, app.address, app.trx, *app.pCreatedOn), nil
}
