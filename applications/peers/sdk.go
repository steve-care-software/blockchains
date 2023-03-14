package peers

import (
	"net/url"
)

// EnterOnConnectFn represents the enter's onConnect func
type EnterOnConnectFn func(peer *url.URL) error

// ExitOnConnectFn represents the exit's onConnect func
type ExitOnConnectFn func(peer *url.URL) error

// Application represents the peers application
type Application interface {
	List() ([]*url.URL, error)
	Connect(peer url.URL) error
}
