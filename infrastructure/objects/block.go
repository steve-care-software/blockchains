package objects

import "time"

// Block represents a block
type Block struct {
	Body  *BlockBody `json:"body"`
	Proof []byte     `json:"proof"`
}

// BlockBody represents the block body
type BlockBody struct {
	Address      []byte    `json:"address"`
	Transactions [][]byte  `json:"transactions"`
	CreatedOn    time.Time `json:"created_on"`
	Parent       []byte    `json:"parent"`
}
