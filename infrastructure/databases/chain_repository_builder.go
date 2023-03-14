package databases

import (
	"errors"

	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainRepositoryBuilder struct {
	hashAdapter            hash.Adapter
	blockRepositoryBuilder blocks.RepositoryBuilder
	database               database_application.Application
	builder                chains.Builder
	pContext               *uint
	pBlockKind             *uint
	pKind                  *uint
}

func createChainRepositoryBuilder(
	hashAdapter hash.Adapter,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	database database_application.Application,
	builder chains.Builder,
) chains.RepositoryBuilder {
	out := chainRepositoryBuilder{
		hashAdapter:            hashAdapter,
		blockRepositoryBuilder: blockRepositoryBuilder,
		database:               database,
		builder:                builder,
		pContext:               nil,
		pBlockKind:             nil,
		pKind:                  nil,
	}

	return &out
}

// Create initializes the builder
func (app *chainRepositoryBuilder) Create() chains.RepositoryBuilder {
	return createChainRepositoryBuilder(
		app.hashAdapter,
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

// WithBlockKind adds a block kind to the builder
func (app *chainRepositoryBuilder) WithBlockKind(blockKind uint) chains.RepositoryBuilder {
	app.pContext = &blockKind
	return app
}

// WithKind adds a kind to the builder
func (app *chainRepositoryBuilder) WithKind(kind uint) chains.RepositoryBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Repository instance
func (app *chainRepositoryBuilder) Now() (chains.Repository, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a chain Repository instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a chain Repository instance")
	}

	builder := app.blockRepositoryBuilder.Create().WithContext(*app.pContext)
	if app.pBlockKind != nil {
		builder.WithKind(*app.pBlockKind)
	}

	blockRepository, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return createChainRepository(
		app.hashAdapter,
		blockRepository,
		app.database,
		app.builder,
		*app.pContext,
		*app.pKind,
	), nil
}
