package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/genesis"
	database_application "github.com/steve-care-software/databases/applications"
)

type genesisRepositoryBuilder struct {
	database database_application.Application
	builder  genesis.Builder
	pContext *uint
}

func createGenesisRepositoryBuilder(
	database database_application.Application,
	builder genesis.Builder,
) genesis.RepositoryBuilder {
	out := genesisRepositoryBuilder{
		database: database,
		builder:  builder,
		pContext: nil,
	}

	return &out
}

// Create initializes the builder
func (app *genesisRepositoryBuilder) Create() genesis.RepositoryBuilder {
	return createGenesisRepositoryBuilder(app.database, app.builder)
}

// WithContext adds a context to the builder
func (app *genesisRepositoryBuilder) WithContext(context uint) genesis.RepositoryBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Repository instance
func (app *genesisRepositoryBuilder) Now() (genesis.Repository, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a genesis Repository instance")
	}

	return createGenesisRepository(
		app.database,
		app.builder,
		*app.pContext,
	), nil
}
