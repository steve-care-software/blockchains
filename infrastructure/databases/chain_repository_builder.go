package databases

import (
	"errors"

	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/genesis"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainRepositoryBuilder struct {
	hashAdapter              hash.Adapter
	genesisRepositoryBuilder genesis.RepositoryBuilder
	blockRepositoryBuilder   blocks.RepositoryBuilder
	database                 database_application.Application
	builder                  chains.Builder
	pContext                 *uint
}

func createChainRepositoryBuilder(
	hashAdapter hash.Adapter,
	genesisRepositoryBuilder genesis.RepositoryBuilder,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	database database_application.Application,
	builder chains.Builder,
) chains.RepositoryBuilder {
	out := chainRepositoryBuilder{
		hashAdapter:              hashAdapter,
		genesisRepositoryBuilder: genesisRepositoryBuilder,
		blockRepositoryBuilder:   blockRepositoryBuilder,
		database:                 database,
		builder:                  builder,
	}

	return &out
}

// Create initializes the builder
func (app *chainRepositoryBuilder) Create() chains.RepositoryBuilder {
	return createChainRepositoryBuilder(
		app.hashAdapter,
		app.genesisRepositoryBuilder,
		app.blockRepositoryBuilder,
		app.database,
		app.builder,
	)
}

// WithContext adds a context to the builder
func (app *chainRepositoryBuilder) WithContext(context uint) chains.RepositoryBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Repository instance
func (app *chainRepositoryBuilder) Now() (chains.Repository, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a chain Repository instance")
	}

	genRepository, err := app.genesisRepositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	blockRepository, err := app.blockRepositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	return createChainRepository(
		app.hashAdapter,
		genRepository,
		blockRepository,
		app.database,
		app.builder,
		*app.pContext,
	), nil
}
