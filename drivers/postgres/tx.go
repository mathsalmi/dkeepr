package postgres

import (
	"database/sql"

	"bitbucket.org/mathsalmi/dkeepr/drivers"
)

// Begin starts a new transaction
func (p *Postgres) Begin() (drivers.Transaction, error) {
	_, err := p.DB().Exec("begin transaction")
	if err != nil {
		return nil, drivers.ErrBegin
	}

	return &Transaction{db: p.DB()}, nil
}

// Transaction represents a transaction of the DB
type Transaction struct {
	err error
	db  *sql.DB
}

// Add func adds a operations to the current transaction
func (tx *Transaction) Add(fn func() error) drivers.Transaction {
	if tx.err == nil {
		tx.err = fn()
	}

	return tx
}

// CommitOrRollback either commits or rollsback the current transaction
func (tx *Transaction) CommitOrRollback() error {
	if tx.err != nil {
		tx.db.Exec("rollback") // TODO: handle err
		return drivers.ErrCommit
	}

	_, err := tx.db.Exec("commit")
	if err != nil {
		return drivers.ErrCommit
	}

	return nil
}
