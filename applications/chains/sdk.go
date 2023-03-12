package chains

import (
	"github.com/steve-care-software/blockchains/applications/chains/blocks"
	"github.com/steve-care-software/blockchains/applications/chains/peers"
	"github.com/steve-care-software/blockchains/applications/chains/transactions"
	"github.com/steve-care-software/blockchains/applications/chains/wallets"
	"github.com/steve-care-software/blockchains/domain/chains"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
)

// Application represents the chain application
type Application interface {
	List() ([]string, error)
	Insert(name string, genesis genesis.Genesis) error
	Retrieve(name string) (chains.Chain, error)
	Block(chain chains.Chain) (blocks.Application, error)
	Transaction(chain chains.Chain) (transactions.Application, error)
	Wallet(chain chains.Chain) (wallets.Application, error)
	Peer(chain chains.Chain) (peers.Application, error)
}
