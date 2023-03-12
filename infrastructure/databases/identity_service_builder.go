package databases

import (
	"errors"

	"github.com/steve-care-software/blockchains/domain/identities"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type identityServiceBuilder struct {
	hashAdapter hash.Adapter
	database    database_application.Application
	repositoryBuilder  identities.RepositoryBuilder
	pContext    *uint
}

func createIdentityServiceBuilder(
	hashAdapter hash.Adapter,
	database database_application.Application,
	repositoryBuilder  identities.RepositoryBuilder,
) identities.ServiceBuilder {
	out := identityServiceBuilder{
		hashAdapter: hashAdapter,
		database:    database,
		repositoryBuilder:  repositoryBuilder,
		pContext:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *identityServiceBuilder) Create() identities.ServiceBuilder {
	return createIdentityServiceBuilder(app.hashAdapter, app.database, app.repositoryBuilder)
}

// WithContext adds a context to the builder
func (app *identityServiceBuilder) WithContext(context uint) identities.ServiceBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Service instance
func (app *identityServiceBuilder) Now() (identities.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a Service instance")
	}

    repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).Now()
    if err != nil {
        return nil, err
    }

	return createIdentityService(
		app.hashAdapter,
		app.database,
		repository,
		*app.pContext,
	), nil
}
