package drivers

import "bitbucket.org/mathsalmi/dkeepr/errors"

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
	Err() error
	Add(fn func() error) Transaction
	CommitOrRollback() error
}

// Tx is the implementation of Transaction interface
type Tx struct {
	err error
}

// Err returns errors during this transaction
func (tx *Tx) Err() error {
	return tx.err
}

// Add func adds a operations to the current transaction
func (tx *Tx) Add(fn func() error) Transaction {
	if tx.err == nil {
		tx.err = fn()
	}

	return tx
}

// CommitOrRollback either commits or rollsback the current transaction
func (tx *Tx) CommitOrRollback() error {
	return errs.ErrMethodNotImpl
}
