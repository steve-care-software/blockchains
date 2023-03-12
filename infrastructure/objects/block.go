package objects

import "time"

// Block represents a block
type Block struct {
	Body  *Body  `json:"body"`
	Proof []byte `json:"proof"`
}

// Body represents the block body
type Body struct {
	Address      []byte    `json:"address"`
	Transactions [][]byte  `json:"transactions"`
	CreatedOn    time.Time `json:"created_on"`
	Parent       []byte    `json:"parent"`
}
