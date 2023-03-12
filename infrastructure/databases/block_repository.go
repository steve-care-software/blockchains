package databases

import (
	"encoding/json"
	"math/big"

	"github.com/steve-care-software/blockchains/applications"
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/transactions"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type blockRepository struct {
	hashAdapter   hash.Adapter
	trxRepository transactions.Repository
	database      database_application.Application
	builder       blocks.Builder
	bodyBuilder   blocks.BodyBuilder
	context       uint
}

func createBlockRepository(
	hashAdapter hash.Adapter,
	trxRepository transactions.Repository,
	database database_application.Application,
	builder blocks.Builder,
	bodyBuilder blocks.BodyBuilder,
	context uint,
) blocks.Repository {
	out := blockRepository{
		hashAdapter:   hashAdapter,
		trxRepository: trxRepository,
		database:      database,
		builder:       builder,
		bodyBuilder:   bodyBuilder,
		context:       context,
	}

	return &out
}

// List returns the list of blocks
func (app *blockRepository) List() ([]hash.Hash, error) {
	contentKeys, err := app.database.ContentKeysByKind(app.context, applications.KindBlock)
	if err != nil {
		return nil, err
	}

	hashes := []hash.Hash{}
	list := contentKeys.List()
	for _, oneKey := range list {
		hashes = append(hashes, oneKey.Hash())
	}

	return hashes, nil
}

// Retrieve retrieves a block by hash
func (app *blockRepository) Retrieve(blockHash hash.Hash) (blocks.Block, error) {
	js, err := app.database.ReadByHash(app.context, blockHash)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Block)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	trxHashes := []hash.Hash{}
	for _, oneTrxBytes := range ins.Body.Transactions {
		pHash, err := app.hashAdapter.FromBytes(oneTrxBytes)
		if err != nil {
			return nil, err
		}

		trxHashes = append(trxHashes, *pHash)
	}

	transactions, err := app.trxRepository.RetrieveList(trxHashes)
	if err != nil {
		return nil, err
	}

	bodyBuilder := app.bodyBuilder.Create().WithAddress(ins.Body.Address).WithTransactions(transactions).CreatedOn(ins.Body.CreatedOn)
	if ins.Body.Parent != nil {
		pParentHash, err := app.hashAdapter.FromBytes(ins.Body.Parent)
		if err != nil {
			return nil, err
		}

		bodyBuilder.WithParent(*pParentHash)
	}

	body, err := bodyBuilder.Now()
	if err != nil {
		return nil, err
	}

	proof := big.NewInt(0).SetBytes(ins.Proof)
	return app.builder.Create().
		WithBody(body).
		WithProof(proof).
		Now()
}
