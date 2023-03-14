package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/transactions"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type blockRepositoryBuilder struct {
	hashAdapter          hash.Adapter
	trxRepositoryBuilder transactions.RepositoryBuilder
	database             database_application.Application
	builder              blocks.Builder
	bodyBuilder          blocks.BodyBuilder
	pContext             *uint
	pKind                *uint
}

func createBlockRepositoryBuilder(
	hashAdapter hash.Adapter,
	trxRepositoryBuilder transactions.RepositoryBuilder,
	database database_application.Application,
	builder blocks.Builder,
	bodyBuilder blocks.BodyBuilder,
) blocks.RepositoryBuilder {
	out := blockRepositoryBuilder{
		hashAdapter:          hashAdapter,
		trxRepositoryBuilder: trxRepositoryBuilder,
		database:             database,
		builder:              builder,
		bodyBuilder:          bodyBuilder,
		pContext:             nil,
		pKind:                nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockRepositoryBuilder) Create() blocks.RepositoryBuilder {
	return createBlockRepositoryBuilder(
		app.hashAdapter,
		app.trxRepositoryBuilder,
		app.database,
		app.builder,
		app.bodyBuilder,
	)
}

// WithContext adds a context to the builder
func (app *blockRepositoryBuilder) WithContext(context uint) blocks.RepositoryBuilder {
	app.pContext = &context
	return app
}

// WithKind adds a kind to the builder
func (app *blockRepositoryBuilder) WithKind(kind uint) blocks.RepositoryBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Repository instance
func (app *blockRepositoryBuilder) Now() (blocks.Repository, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a block Repository instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a block Repository instance")
	}

	trxRepository, err := app.trxRepositoryBuilder.Create().WithContext(*app.pContext).WithKind(*app.pKind).Now()
	if err != nil {
		return nil, err
	}

	return createBlockRepository(
		app.hashAdapter,
		trxRepository,
		app.database,
		app.builder,
		app.bodyBuilder,
		*app.pContext,
		*app.pKind,
	), nil
}
