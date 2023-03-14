package databases

import (
	"encoding/json"
	"fmt"

	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainRepository struct {
	hashAdapter     hash.Adapter
	blockRepository blocks.Repository
	database        database_application.Application
	builder         chains.Builder
	context         uint
	kind            uint
}

func createChainRepository(
	hashAdapter hash.Adapter,
	blockRepository blocks.Repository,
	database database_application.Application,
	builder chains.Builder,
	context uint,
	kind uint,
) chains.Repository {
	out := chainRepository{
		hashAdapter:     hashAdapter,
		blockRepository: blockRepository,
		database:        database,
		builder:         builder,
		context:         context,
		kind:            kind,
	}

	return &out
}

// List returns the chain names
func (app *chainRepository) List() ([]string, error) {
	pHash, err := app.hashAdapter.FromBytes([]byte(chainList))
	if err != nil {
		return nil, err
	}

	js, err := app.database.ReadByHash(app.context, app.kind, *pHash)
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

	js, err := app.database.ReadByHash(app.context, app.kind, *pHash)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Chain)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithName(ins.Name).WithRoot(ins.Root)
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
