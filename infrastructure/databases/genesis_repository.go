package databases

import (
	"encoding/json"

	"github.com/steve-care-software/blockchains/domain/genesis"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type genesisRepository struct {
	database database_application.Application
	builder  genesis.Builder
	context  uint
}

func createGenesisRepository(
	database database_application.Application,
	builder genesis.Builder,
	context uint,
) genesis.Repository {
	out := genesisRepository{
		database: database,
		builder:  builder,
		context:  context,
	}

	return &out
}

// Retrieve retrieves a genesis by hash
func (app *genesisRepository) Retrieve(hash hash.Hash) (genesis.Genesis, error) {
	js, err := app.database.ReadByHash(app.context, hash)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Genesis)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithDifficulty(ins.Difficulty).
		WithReward(ins.Reward).
		WithHalving(ins.Halving).
		WithMiningValue(ins.MiningValue).
		WithMaxTrxPerBlock(ins.MaxTrxPerBlock).
		WithBlockDuration(ins.BlockDuration).
		CreatedOn(ins.CreatedOn).
		Now()
}
