package databases

import (
	"encoding/json"
	"errors"
	"fmt"

	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/blockchains/infrastructure/objects"
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chainService struct {
	hashAdapter     hash.Adapter
	repository      chains.Repository
	blockRepository blocks.Repository
	database        database_application.Application
	context         uint
	kind            uint
}

func createChainService(
	hashAdapter hash.Adapter,
	repository chains.Repository,
	blockRepository blocks.Repository,
	database database_application.Application,
	context uint,
	kind uint,
) chains.Service {
	out := chainService{
		hashAdapter:     hashAdapter,
		repository:      repository,
		blockRepository: blockRepository,
		database:        database,
		context:         context,
		kind:            kind,
	}

	return &out
}

// Insert inserts a chain
func (app *chainService) Insert(chain chains.Chain) error {
	name := chain.Name()
	_, err := app.repository.Retrieve(name)
	if err == nil {
		str := fmt.Sprintf("the Chain (name: %s) already exists", name)
		return errors.New(str)
	}

	var headBytes []byte
	if chain.HasHead() {
		blockHash := chain.Head().Hash()
		_, err = app.blockRepository.Retrieve(blockHash)
		if err != nil {
			str := fmt.Sprintf("the chain (name: %s) cannot be inserted because it contains a head block (hash: %s) that could not be retrieved: %s", name, blockHash.String(), err.Error())
			return errors.New(str)
		}

		headBytes = blockHash.Bytes()
	}

	ins := objects.Chain{
		Name: chain.Name(),
		Root: chain.Root(),
		Head: headBytes,
	}

	js, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	pHash, err := app.hashAdapter.FromBytes([]byte(fmt.Sprintf(chainByName, name)))
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		app.kind,
		*pHash,
		js,
	)
}
