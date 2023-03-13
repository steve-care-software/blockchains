package databases

import (
	"errors"

	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/genesis"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainServiceBuilder struct {
	hashAdapter              hash.Adapter
	repositoryBuilder        chains.RepositoryBuilder
	blockRepositoryBuilder   blocks.RepositoryBuilder
	genesisRepositoryBuilder genesis.RepositoryBuilder
	genesisServiceBuilder    genesis.ServiceBuilder
	database                 database_application.Application
	pContext                 *uint
}

func createChainServiceBuilder(
	hashAdapter hash.Adapter,
	repositoryBuilder chains.RepositoryBuilder,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	genesisRepositoryBuilder genesis.RepositoryBuilder,
	genesisServiceBuilder genesis.ServiceBuilder,
	database database_application.Application,
) chains.ServiceBuilder {
	out := chainServiceBuilder{
		hashAdapter:              hashAdapter,
		repositoryBuilder:        repositoryBuilder,
		blockRepositoryBuilder:   blockRepositoryBuilder,
		genesisRepositoryBuilder: genesisRepositoryBuilder,
		genesisServiceBuilder:    genesisServiceBuilder,
		database:                 database,
		pContext:                 nil,
	}

	return &out
}

// Create initializes the builder
func (app *chainServiceBuilder) Create() chains.ServiceBuilder {
	return createChainServiceBuilder(
		app.hashAdapter,
		app.repositoryBuilder,
		app.blockRepositoryBuilder,
		app.genesisRepositoryBuilder,
		app.genesisServiceBuilder,
		app.database,
	)
}

// WithContext adds a context to the builder
func (app *chainServiceBuilder) WithContext(context uint) chains.ServiceBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Service instance
func (app *chainServiceBuilder) Now() (chains.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a chain Service instance")
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	blockRepository, err := app.blockRepositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	genesisRepository, err := app.genesisRepositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	genesisService, err := app.genesisServiceBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	return createChainService(
		app.hashAdapter,
		repository,
		blockRepository,
		genesisRepository,
		genesisService,
		app.database,
		*app.pContext,
	), nil
}
