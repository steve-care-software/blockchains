package databases

import (
	application_identity "github.com/steve-care-software/blockchains/applications/identities"
	"github.com/steve-care-software/blockchains/domain/chains"
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
	"github.com/steve-care-software/blockchains/domain/chains/transactions"
	"github.com/steve-care-software/blockchains/domain/identities"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

const identityList = "identity:list"
const identityListDeleted = "identity:list:deleted"

const chainList = "chain:list"
const chainByName = "chain:by_name:%s"

// NewChainServiceBuilder creates a new chain service builder
func NewChainServiceBuilder(
	repositoryBuilder chains.RepositoryBuilder,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	genesisRepositoryBuilder genesis.RepositoryBuilder,
	genesisServiceBuilder genesis.ServiceBuilder,
	database database_application.Application,
) chains.ServiceBuilder {
	hashAdapter := hash.NewAdapter()
	return createChainServiceBuilder(
		hashAdapter,
		repositoryBuilder,
		blockRepositoryBuilder,
		genesisRepositoryBuilder,
		genesisServiceBuilder,
		database,
	)
}

// NewChainRepositoryBuilder creates a new chain repository builder
func NewChainRepositoryBuilder(
	genesisRepositoryBuilder genesis.RepositoryBuilder,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	database database_application.Application,
) chains.RepositoryBuilder {
	hashAdapter := hash.NewAdapter()
	builder := chains.NewBuilder()
	return createChainRepositoryBuilder(
		hashAdapter,
		genesisRepositoryBuilder,
		blockRepositoryBuilder,
		database,
		builder,
	)
}

// NewBlockServiceBuilder creates a new block service builder
func NewBlockServiceBuilder(
	database database_application.Application,
	repository blocks.Repository,
	trxService transactions.Service,
) blocks.ServiceBuilder {
	return createBlockServiceBuilder(
		database,
		repository,
		trxService,
	)
}

// NewBlockRepositoryBuilder creates a new block repository builder
func NewBlockRepositoryBuilder(
	trxRepositoryBuilder transactions.RepositoryBuilder,
	database database_application.Application,
) blocks.RepositoryBuilder {
	hashAdapter := hash.NewAdapter()
	builder := blocks.NewBuilder()
	bodyBuilder := blocks.NewBodyBuilder()
	return createBlockRepositoryBuilder(
		hashAdapter,
		trxRepositoryBuilder,
		database,
		builder,
		bodyBuilder,
	)
}

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
