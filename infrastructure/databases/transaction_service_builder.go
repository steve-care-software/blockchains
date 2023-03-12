package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/chains/transactions"
	database_application "github.com/steve-care-software/databases/applications"
)

type transactionServiceBuilder struct {
	database          database_application.Application
	repositoryBuilder transactions.RepositoryBuilder
	pContext          *uint
	pKind             *uint
}

func createTransactionServiceBuilder(
	database database_application.Application,
	repositoryBuilder transactions.RepositoryBuilder,
) transactions.ServiceBuilder {
	out := transactionServiceBuilder{
		database:          database,
		repositoryBuilder: repositoryBuilder,
		pContext:          nil,
		pKind:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionServiceBuilder) Create() transactions.ServiceBuilder {
	return createTransactionServiceBuilder(
		app.database,
		app.repositoryBuilder,
	)
}

// WithContext adds a context to the builder
func (app *transactionServiceBuilder) WithContext(context uint) transactions.ServiceBuilder {
	app.pContext = &context
	return app
}

// WithKind adds a kind to the builder
func (app *transactionServiceBuilder) WithKind(kind uint) transactions.ServiceBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Service instance
func (app *transactionServiceBuilder) Now() (transactions.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a transaction Service")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a transaction Service")
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).WithKind(*app.pKind).Now()
	if err != nil {
		return nil, err
	}

	return createTransactionService(
		app.database,
		repository,
		*app.pContext,
		*app.pKind,
	), nil
}
