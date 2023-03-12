package daemons

// Daemon represents the blockchain daemon
type Daemon interface {
	Start()
	Stop() error
}
