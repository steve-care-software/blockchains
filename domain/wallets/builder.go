package wallets

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	address     []byte
	amount      uint
	trx         transactions.Transactions
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		address:     nil,
		amount:      0,
		trx:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithAddress adds an address to the builder
func (app *builder) WithAddress(address []byte) Builder {
	app.address = address
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint) Builder {
	app.amount = amount
	return app
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(trx transactions.Transactions) Builder {
	app.trx = trx
	return app
}

// Now builds a new Wallet instance
func (app *builder) Now() (Wallet, error) {
	if app.address != nil && len(app.address) <= 0 {
		app.address = nil
	}

	if app.address == nil {
		return nil, errors.New("the address is mandatory in order to build a Wallet instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Wallet instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions are mandatory in order to build a Wallet instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.address,
		[]byte(fmt.Sprintf("%d", app.amount)),
		app.trx.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createWallet(*pHash, app.address, app.amount, app.trx), nil
}
