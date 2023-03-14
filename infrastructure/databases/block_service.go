package databases

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/domain/transactions"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
)

type blockService struct {
	database   database_application.Application
	repository blocks.Repository
	trxService transactions.Service
	context    uint
	kind       uint
}

func createBlockService(
	database database_application.Application,
	repository blocks.Repository,
	trxService transactions.Service,
	context uint,
	kind uint,
) blocks.Service {
	out := blockService{
		database:   database,
		repository: repository,
		trxService: trxService,
		context:    context,
		kind:       kind,
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
		Body: &objects.BlockBody{
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
		app.kind,
		hash,
		js,
	)
}
