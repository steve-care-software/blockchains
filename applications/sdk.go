package applications

import (
	"github.com/steve-care-software/blockchains/applications/chains"
)

const (
	// KindIdentities represents an identities kind
	KindIdentities (uint) = iota

	// KindGenesis represents a genesis kind
	KindGenesis

	// KindBlock represents a block kind
	KindBlock

	// KindChain represents a chain kind
	KindChain

	// KindTransaction represents a transaction kind
	KindTransaction
)

// Application represents the blockchain application
type Application interface {
	Chain() chains.Application
}
