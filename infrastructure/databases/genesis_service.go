package databases

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/steve-care-software/blockchains/applications"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
)

type genesisService struct {
	database   database_application.Application
	repository genesis.Repository
	context    uint
}

func createGenesisService(
	database database_application.Application,
	repository genesis.Repository,
	context uint,
) genesis.Service {
	out := genesisService{
		database:   database,
		repository: repository,
		context:    context,
	}

	return &out
}

// Insert inserts a genesis instance
func (app *genesisService) Insert(genesis genesis.Genesis) error {
	hash := genesis.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		str := fmt.Sprintf("the Genesis (hash: %s) already exists", hash.String())
		return errors.New(str)
	}

	ins := objects.Genesis{
		Hash:           hash.String(),
		Difficulty:     genesis.Difficulty(),
		Reward:         genesis.Reward(),
		Halving:        genesis.Halving(),
		MiningValue:    genesis.MiningValue(),
		MaxTrxPerBlock: genesis.MaxTrxPerBlock(),
		BlockDuration:  genesis.BlockDuration(),
		CreatedOn:      genesis.CreatedOn(),
	}

	js, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		applications.KindIdentity,
		hash,
		js,
	)
}
