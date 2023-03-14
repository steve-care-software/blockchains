package databases

import (
	"errors"

	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainServiceBuilder struct {
	hashAdapter            hash.Adapter
	repositoryBuilder      chains.RepositoryBuilder
	blockRepositoryBuilder blocks.RepositoryBuilder
	database               database_application.Application
	pContext               *uint
	pKind                  *uint
}

func createChainServiceBuilder(
	hashAdapter hash.Adapter,
	repositoryBuilder chains.RepositoryBuilder,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	database database_application.Application,
) chains.ServiceBuilder {
	out := chainServiceBuilder{
		hashAdapter:            hashAdapter,
		repositoryBuilder:      repositoryBuilder,
		blockRepositoryBuilder: blockRepositoryBuilder,
		database:               database,
		pContext:               nil,
		pKind:                  nil,
	}

	return &out
}

// Create initializes the builder
func (app *chainServiceBuilder) Create() chains.ServiceBuilder {
	return createChainServiceBuilder(
		app.hashAdapter,
		app.repositoryBuilder,
		app.blockRepositoryBuilder,
		app.database,
	)
}

// WithContext adds a context to the builder
func (app *chainServiceBuilder) WithContext(context uint) chains.ServiceBuilder {
	app.pContext = &context
	return app
}

// WithKind adds a kind to the builder
func (app *chainServiceBuilder) WithKind(kind uint) chains.ServiceBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Service instance
func (app *chainServiceBuilder) Now() (chains.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a chain Service instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a chain Service instance")
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	blockRepository, err := app.blockRepositoryBuilder.Create().WithContext(*app.pContext).Now()
	if err != nil {
		return nil, err
	}

	return createChainService(
		app.hashAdapter,
		repository,
		blockRepository,
		app.database,
		*app.pContext,
		*app.pKind,
	), nil
}
