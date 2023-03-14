package applications

import (
	"github.com/steve-care-software/blockchains/applications/blocks"
	"github.com/steve-care-software/blockchains/applications/peers"
	"github.com/steve-care-software/blockchains/applications/transactions"
	chains "github.com/steve-care-software/blockchains/domain"
)

// Application represents the chain application
type Application interface {
	List() ([]string, error)
	Insert(name string, root []byte) error
	Retrieve(name string) (chains.Chain, error)
	Block(chain chains.Chain) (blocks.Application, error)
	Transaction(chain chains.Chain) (transactions.Application, error)
	Peer(chain chains.Chain) (peers.Application, error)
}
