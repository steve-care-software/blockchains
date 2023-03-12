package databases

import (
	"encoding/json"
	"fmt"

	"github.com/steve-care-software/blockchains/domain/chains"
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainRepository struct {
	hashAdapter       hash.Adapter
	genesisRepository genesis.Repository
	blockRepository   blocks.Repository
	database          database_application.Application
	builder           chains.Builder
	context           uint
}

func createChainRepository(
	hashAdapter hash.Adapter,
	genesisRepository genesis.Repository,
	blockRepository blocks.Repository,
	database database_application.Application,
	builder chains.Builder,
	context uint,
) chains.Repository {
	out := chainRepository{
		hashAdapter:       hashAdapter,
		genesisRepository: genesisRepository,
		blockRepository:   blockRepository,
		database:          database,
		builder:           builder,
		context:           context,
	}

	return &out
}

// List returns the chain names
func (app *chainRepository) List() ([]string, error) {
	pHash, err := app.hashAdapter.FromBytes([]byte(chainList))
	if err != nil {
		return nil, err
	}

	js, err := app.database.ReadByHash(app.context, *pHash)
	if err != nil {
		return nil, err
	}

	list := new([]string)
	err = json.Unmarshal(js, list)
	if err != nil {
		return nil, err
	}

	return *list, nil
}

// Retrieve retrieves a chain by name
func (app *chainRepository) Retrieve(name string) (chains.Chain, error) {
	pHash, err := app.hashAdapter.FromBytes([]byte(fmt.Sprintf(chainByName, name)))
	if err != nil {
		return nil, err
	}

	js, err := app.database.ReadByHash(app.context, *pHash)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Chain)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	pGenesisHash, err := app.hashAdapter.FromBytes(ins.Root)
	if err != nil {
		return nil, err
	}

	genesis, err := app.genesisRepository.Retrieve(*pGenesisHash)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithName(ins.Name).WithRoot(genesis)
	if ins.Head != nil {
		pBlockHash, err := app.hashAdapter.FromBytes(ins.Head)
		if err != nil {
			return nil, err
		}

		block, err := app.blockRepository.Retrieve(*pBlockHash)
		if err != nil {
			return nil, err
		}

		builder.WithHead(block)
	}

	return builder.Now()
}
