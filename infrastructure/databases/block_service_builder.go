package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/transactions"
	database_application "github.com/steve-care-software/databases/applications"
)

type blockServiceBuilder struct {
	database   database_application.Application
	repository blocks.Repository
	trxService transactions.Service
	pContext   *uint
}

func createBlockServiceBuilder(
	database database_application.Application,
	repository blocks.Repository,
	trxService transactions.Service,
) blocks.ServiceBuilder {
	out := blockServiceBuilder{
		database:   database,
		repository: repository,
		trxService: trxService,
		pContext:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockServiceBuilder) Create() blocks.ServiceBuilder {
	return createBlockServiceBuilder(
		app.database,
		app.repository,
		app.trxService,
	)
}

// WithContext adds a context to the builder
func (app *blockServiceBuilder) WithContext(context uint) blocks.ServiceBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Service instance
func (app *blockServiceBuilder) Now() (blocks.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a block Service instance")
	}

	return createBlockService(
		app.database,
		app.repository,
		app.trxService,
		*app.pContext,
	), nil
}
