package drivers

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
