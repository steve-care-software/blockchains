package databases

import (
	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/transactions"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

const identityList = "identity:list"
const identityListDeleted = "identity:list:deleted"

const chainList = "chain:list"
const chainByName = "chain:by_name:%s"

// NewTransactionServiceBuilder creates a new transaction service builder
func NewTransactionServiceBuilder(
	database database_application.Application,
	repositoryBuilder transactions.RepositoryBuilder,
) transactions.ServiceBuilder {
	return createTransactionServiceBuilder(
		database,
		repositoryBuilder,
	)
}

// NewTransactionRepositoryBuilder creates a new transaction repository builder
func NewTransactionRepositoryBuilder(
	database database_application.Application,
	trxBuilder transactions.TransactionBuilder,
) transactions.RepositoryBuilder {
	hashAdapter := hash.NewAdapter()
	builder := transactions.NewBuilder()
	bodyBuilder := transactions.NewBodyBuilder()
	return createTransactionRepositoryBuilder(
		hashAdapter,
		database,
		builder,
		trxBuilder,
		bodyBuilder,
	)
}

// NewChainServiceBuilder creates a new chain service builder
func NewChainServiceBuilder(
	repositoryBuilder chains.RepositoryBuilder,
	blockRepositoryBuilder blocks.RepositoryBuilder,
	database database_application.Application,
) chains.ServiceBuilder {
	hashAdapter := hash.NewAdapter()
	return createChainServiceBuilder(
		hashAdapter,
		repositoryBuilder,
		blockRepositoryBuilder,
		database,
	)
}

// NewChainRepositoryBuilder creates a new chain repository builder
func NewChainRepositoryBuilder(
	blockRepositoryBuilder blocks.RepositoryBuilder,
	database database_application.Application,
) chains.RepositoryBuilder {
	hashAdapter := hash.NewAdapter()
	builder := chains.NewBuilder()
	return createChainRepositoryBuilder(
		hashAdapter,
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
