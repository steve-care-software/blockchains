package peers

import (
	"net/url"
)

// Application represents the peers application
type Application interface {
	List() ([]*url.URL, error)
	Connect(peer url.URL) error
}
