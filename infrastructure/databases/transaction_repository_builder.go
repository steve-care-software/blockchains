package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/transactions"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionRepositoryBuilder struct {
	hashAdapter hash.Adapter
	database    database_application.Application
	builder     transactions.Builder
	trxBuilder  transactions.TransactionBuilder
	bodyBuilder transactions.BodyBuilder
	pContext    *uint
	pKind       *uint
}

func createTransactionRepositoryBuilder(
	hashAdapter hash.Adapter,
	database database_application.Application,
	builder transactions.Builder,
	trxBuilder transactions.TransactionBuilder,
	bodyBuilder transactions.BodyBuilder,
) transactions.RepositoryBuilder {
	out := transactionRepositoryBuilder{
		hashAdapter: hashAdapter,
		database:    database,
		builder:     builder,
		trxBuilder:  trxBuilder,
		bodyBuilder: bodyBuilder,
	}

	return &out
}

// Create initializes the builder
func (app *transactionRepositoryBuilder) Create() transactions.RepositoryBuilder {
	return createTransactionRepositoryBuilder(
		app.hashAdapter,
		app.database,
		app.builder,
		app.trxBuilder,
		app.bodyBuilder,
	)
}

// WithContext adds a context to the builder
func (app *transactionRepositoryBuilder) WithContext(context uint) transactions.RepositoryBuilder {
	app.pContext = &context
	return app
}

// WithKind adds a kind to the builder
func (app *transactionRepositoryBuilder) WithKind(kind uint) transactions.RepositoryBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Repository instance
func (app *transactionRepositoryBuilder) Now() (transactions.Repository, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a transaction's Repository instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a transaction's Repository instance")
	}

	return createTransactionRepository(
		app.hashAdapter,
		app.database,
		app.builder,
		app.trxBuilder,
		app.bodyBuilder,
		*app.pContext,
		*app.pKind,
	), nil
}
