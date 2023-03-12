package databases

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/steve-care-software/blockchains/applications"
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/transactions"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
)

type blockService struct {
	database   database_application.Application
	repository blocks.Repository
	trxService transactions.Service
	context    uint
}

func createBlockService(
	database database_application.Application,
	repository blocks.Repository,
	trxService transactions.Service,
	context uint,
) blocks.Service {
	out := blockService{
		database:   database,
		repository: repository,
		trxService: trxService,
		context:    context,
	}

	return &out
}

// Insert inserts a block
func (app *blockService) Insert(block blocks.Block) error {
	hash := block.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		str := fmt.Sprintf("the Block (hash: %s) already exists", hash.String())
		return errors.New(str)
	}

	body := block.Body()
	var parent []byte
	if body.HasParent() {
		parent = body.Parent().Bytes()
	}

	ins := objects.Block{
		Body: &objects.Body{
			Address:      body.Address(),
			Transactions: nil,
			CreatedOn:    body.CreatedOn(),
			Parent:       parent,
		},
		Proof: block.Proof().Bytes(),
	}

	js, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		applications.KindBlock,
		hash,
		js,
	)
}
