package databases

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionService struct {
	database   database_application.Application
	repository transactions.Repository
	context    uint
	kind       uint
}

func createTransactionService(
	database database_application.Application,
	repository transactions.Repository,
	context uint,
	kind uint,
) transactions.Service {
	out := transactionService{
		database:   database,
		repository: repository,
		context:    context,
		kind:       kind,
	}

	return &out
}

// Insert inserts a transaction
func (app *transactionService) Insert(trx transactions.Transaction) error {
	hash := trx.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		str := fmt.Sprintf("the Transaction (hash: %s) already exists", hash.String())
		return errors.New(str)
	}

	body := trx.Body()
	ins := objects.Transaction{
		Body: &objects.TransactionBody{
			Address:   body.Address().Bytes(),
			Fees:      body.Fees(),
			Reference: body.Reference().Bytes(),
		},
		Signature: trx.Signature(),
		PublicKey: trx.PublicKey(),
	}

	js, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		app.kind,
		hash,
		js,
	)
}

// InsertList inserts a list of transactions
func (app *transactionService) InsertList(list []transactions.Transaction) error {
	for _, oneTrx := range list {
		err := app.Insert(oneTrx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Erase erases a transaction by hash
func (app *transactionService) Erase(hash hash.Hash) error {
	return app.database.EraseByHash(app.context, app.kind, hash)
}

// EraseAll erases a list of transactions by hashes
func (app *transactionService) EraseAll(hashes []hash.Hash) error {
	return app.database.EraseAllByHashes(app.context, app.kind, hashes)
}
