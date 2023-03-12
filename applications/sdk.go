package applications

import (
	"github.com/steve-care-software/blockchains/applications/chains"
	"github.com/steve-care-software/blockchains/applications/identities"
)

const (
	// KindIdentities represents an identities kind
	KindIdentities (uint) = iota

	// KindIdentity represents an identity kind
	KindIdentity
)

// Application represents the blockchain application
type Application interface {
	Chain() chains.Application
	Identity() identities.Application
}
