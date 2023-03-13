package objects

// Transaction represents a transaction
type Transaction struct {
	Body      *TransactionBody `json:"body"`
	Signature []byte           `json:"signature"`
}

// TransactionBody represents a transaction body
type TransactionBody struct {
	Address   []byte `json:"address"`
	Fees      uint   `json:"fees"`
	Reference []byte `json:"reference"`
}
