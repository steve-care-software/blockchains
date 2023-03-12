package objects

// Chain represents a chain
type Chain struct {
	Name string `json:"name"`
	Root []byte `json:"root"`
	Head []byte `json:"head"`
}
