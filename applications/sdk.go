package applications

import (
	"github.com/steve-care-software/blockchains/applications/blocks"
	"github.com/steve-care-software/blockchains/applications/peers"
	"github.com/steve-care-software/blockchains/applications/transactions"
	"github.com/steve-care-software/blockchains/applications/wallets"
	"github.com/steve-care-software/blockchains/domain/chains"
)

// Application represents the chain application
type Application interface {
	List() ([]string, error)
	Retrieve(name string) (chains.Chain, error)
	Block(chain chains.Chain) (blocks.Application, error)
	Transaction(chain chains.Chain) (transactions.Application, error)
	Wallet(chain chains.Chain) (wallets.Application, error)
	Peer(chain chains.Chain) (peers.Application, error)
}
