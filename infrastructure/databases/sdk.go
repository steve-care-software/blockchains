package databases

import (
	application_identity "github.com/steve-care-software/blockchains/applications/identities"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
	"github.com/steve-care-software/blockchains/domain/identities"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

const identityList = "identity:list"
const identityListDeleted = "identity:list:deleted"

// NewGenesisServiceBuilder creates a new genesis service builder
func NewGenesisServiceBuilder(
	database database_application.Application,
	repositoryBuilder genesis.RepositoryBuilder,
) genesis.ServiceBuilder {
	return createGenesisServiceBuilder(database, repositoryBuilder)
}

// NewGenesisRepositoryBuilder creates a new genesis repository builder
func NewGenesisRepositoryBuilder(
	database database_application.Application,
) genesis.RepositoryBuilder {
	builder := genesis.NewBuilder()
	return createGenesisRepositoryBuilder(database, builder)
}

// NewIdentityApplication creates a new identity application
func NewIdentityApplication(
	repositoryBuilder identities.RepositoryBuilder,
	serviceBuilder identities.ServiceBuilder,
	dbName string,
	database database_application.Application,
) application_identity.Application {
	return createIdentityApplication(database, repositoryBuilder, serviceBuilder, dbName)
}

// NewIdentityServiceBuilder creates a new identity service builder
func NewIdentityServiceBuilder(
	repositoryBuilder identities.RepositoryBuilder,
	database database_application.Application,
) identities.ServiceBuilder {
	hashAdapter := hash.NewAdapter()
	return createIdentityServiceBuilder(hashAdapter, database, repositoryBuilder)
}

// NewIdentityRepositoryBuilder creates a new identy repository builder
func NewIdentityRepositoryBuilder(
	database database_application.Application,
) identities.RepositoryBuilder {
	hashAdapter := hash.NewAdapter()
	builder := identities.NewBuilder()
	return createIdentityRepositoryBuilder(hashAdapter, database, builder)
}
