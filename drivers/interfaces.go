package drivers

import "errors"

// Errors
var (
	ErrNoResult = errors.New("ORM: no result")
)

// Entity represents a parsed entity
type Entity interface {
	Tablename() string

	Pk() []string
	Pkvalues() []interface{}

	Columns() []string
	Values() []interface{}
}
