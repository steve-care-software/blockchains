package identities

import "time"

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPrivate(private []byte) Builder
	WithPublic(public []byte) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Private() []byte
	Public() []byte
	CreatedOn() time.Time
}
