package domain

import (
	"github.com/steve-care-software/blockchains/domain/blocks"
)

type chain struct {
	name string
	root []byte
	head blocks.Block
}

func createChain(
	name string,
	root []byte,
) Chain {
	return createChainInternally(name, root, nil)
}

func createChainWithHead(
	name string,
	root []byte,
	head blocks.Block,
) Chain {
	return createChainInternally(name, root, head)
}

func createChainInternally(
	name string,
	root []byte,
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
func (obj *chain) Root() []byte {
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
