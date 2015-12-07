package drivers

import "errors"

// Errors
var (
	ErrNoResult = errors.New("ORM: no result")
	ErrNoPk     = errors.New("ORM: obligatory IDs not given")

	// transaction
	ErrBegin  = errors.New("ORM: error opening transaction")
	ErrCommit = errors.New("ORM: there was a problem with something. Rolling back transaction... ")
)

// Entity represents a parsed entity
type Entity interface {
	Tablename() string

	Pk() []string
	Pkvalues() []interface{}

	Columns() []string
	Values() []interface{}
}

// Transaction is the representation of a transaction
type Transaction interface {
	Add(fn func() error) Transaction
	CommitOrRollback() error
}
