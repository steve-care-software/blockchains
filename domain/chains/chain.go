package chains

import (
	"github.com/steve-care-software/blockchains/domain/chains/blocks"
	"github.com/steve-care-software/blockchains/domain/chains/genesis"
)

type chain struct {
	name string
	root genesis.Genesis
	head blocks.Block
}

func createChain(
	name string,
	root genesis.Genesis,
) Chain {
	return createChainInternally(name, root, nil)
}

func createChainWithHead(
	name string,
	root genesis.Genesis,
	head blocks.Block,
) Chain {
	return createChainInternally(name, root, head)
}

func createChainInternally(
	name string,
	root genesis.Genesis,
	head blocks.Block,
) Chain {
	out := chain{
		name: name,
		root: root,
		head: head,
	}

	return &out
}

// Name returns the name
func (obj *chain) Name() string {
	return obj.name
}

// Root returns the root
func (obj *chain) Root() genesis.Genesis {
	return obj.root
}

// HasHead returns true if there is a head, false otherwise
func (obj *chain) HasHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *chain) Head() blocks.Block {
	return obj.head
}