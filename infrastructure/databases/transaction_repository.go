package databases

import (
	"encoding/json"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionRepository struct {
	hashAdapter hash.Adapter
	database    database_application.Application
	builder     transactions.Builder
	trxBuilder  transactions.TransactionBuilder
	bodyBuilder transactions.BodyBuilder
	context     uint
	kind        uint
}

func createTransactionRepository(
	hashAdapter hash.Adapter,
	database database_application.Application,
	builder transactions.Builder,
	trxBuilder transactions.TransactionBuilder,
	bodyBuilder transactions.BodyBuilder,
	context uint,
	kind uint,
) transactions.Repository {
	out := transactionRepository{
		hashAdapter: hashAdapter,
		database:    database,
		builder:     builder,
		trxBuilder:  trxBuilder,
		bodyBuilder: bodyBuilder,
		context:     context,
		kind:        kind,
	}

	return &out
}

// List returns the transaction list
func (app *transactionRepository) List() []hash.Hash {
	contentKeys, err := app.database.ContentKeysByKind(app.context, app.kind)
	if err != nil {
		return []hash.Hash{}
	}

	hashes := []hash.Hash{}
	list := contentKeys.List()
	for _, oneKey := range list {
		hashes = append(hashes, oneKey.Hash())
	}

	return hashes
}

// Retrieve returns transaction by hash
func (app *transactionRepository) Retrieve(trxHash hash.Hash) (transactions.Transaction, error) {
	js, err := app.database.ReadByHash(app.context, app.kind, trxHash)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Transaction)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	pReferenceHash, err := app.hashAdapter.FromBytes(ins.Body.Reference)
	if err != nil {
		return nil, err
	}

	pAddressHash, err := app.hashAdapter.FromBytes(ins.Body.Address)
	if err != nil {
		return nil, err
	}

	body, err := app.bodyBuilder.Create().
		WithAddress(*pAddressHash).
		WithFees(ins.Body.Fees).
		WithReference(*pReferenceHash).
		Now()

	if err != nil {
		return nil, err
	}

	return app.trxBuilder.Create().
		WithBody(body).
		WithSignature(ins.Signature).
		WithPublicKey(ins.PublicKey).
		Now()
}

// RetrieveList retrieves list of transactions by hashes
func (app *transactionRepository) RetrieveList(hashes []hash.Hash) (transactions.Transactions, error) {
	list := []transactions.Transaction{}
	for _, oneHash := range hashes {
		trx, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		list = append(list, trx)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}
