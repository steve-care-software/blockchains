package wallets

import (
	chains "github.com/steve-care-software/blockchains/domain"
	"github.com/steve-care-software/blockchains/domain/wallets"
)

// Builder represents the wallet's application builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	Now() (Application, error)
}

// Application represents the wallet application
type Application interface {
	Retrieve(address []byte) (wallets.Wallet, error)
}
