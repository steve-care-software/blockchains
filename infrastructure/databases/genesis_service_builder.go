package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/chains/genesis"
	database_application "github.com/steve-care-software/databases/applications"
)

type genesisServiceBuilder struct {
	database          database_application.Application
	repositoryBuilder genesis.RepositoryBuilder
	pContext          *uint
}

func createGenesisServiceBuilder(
	database database_application.Application,
	repositoryBuilder genesis.RepositoryBuilder,
) genesis.ServiceBuilder {
	out := genesisServiceBuilder{
		database:          database,
		repositoryBuilder: repositoryBuilder,
		pContext:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *genesisServiceBuilder) Create() genesis.ServiceBuilder {
	return createGenesisServiceBuilder(app.database, app.repositoryBuilder)
}

// WithContext adds a context to the builder
func (app *genesisServiceBuilder) WithContext(context uint) genesis.ServiceBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Service instance
func (app *genesisServiceBuilder) Now() (genesis.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a genesis Service instance")
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	return createGenesisService(
		app.database,
		repository,
		*app.pContext,
	), nil
}
