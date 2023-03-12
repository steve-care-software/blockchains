package applications

import (
	"github.com/steve-care-software/blockchains/applications/chains"
	"github.com/steve-care-software/blockchains/applications/identities"
)

// Application represents the blockchain application
type Application interface {
	Chain() chains.Application
	Identity() identities.Application
}
