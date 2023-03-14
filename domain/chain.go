package domain

import (
	"github.com/steve-care-software/blockchains/domain/blocks"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type chain struct {
	name string
	root hash.Hash
	head blocks.Block
}

func createChain(
	name string,
	root hash.Hash,
) Chain {
	return createChainInternally(name, root, nil)
}

func createChainWithHead(
	name string,
	root hash.Hash,
	head blocks.Block,
) Chain {
	return createChainInternally(name, root, head)
}

func createChainInternally(
	name string,
	root hash.Hash,
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
func (obj *chain) Root() hash.Hash {
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
