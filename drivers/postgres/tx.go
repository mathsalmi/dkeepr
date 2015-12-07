package postgres

import (
	"database/sql"

	"bitbucket.org/mathsalmi/dkeepr/drivers"
	"bitbucket.org/mathsalmi/dkeepr/errors"
)

// Begin starts a new transaction
func (p *Postgres) Begin() (drivers.Transaction, error) {
	_, err := p.DB().Exec("begin transaction")
	if err != nil {
		return nil, errs.ErrBegin
	}

	return &Transaction{db: p.DB()}, nil
}

// Transaction represents a transaction of the DB
type Transaction struct {
	drivers.Tx
	db *sql.DB
}

// CommitOrRollback either commits or rollsback the current transaction
func (tx *Transaction) CommitOrRollback() error {
	if tx.Err() != nil {
		tx.db.Exec("rollback") // TODO: handle err
		return errs.ErrCommit
	}

	_, err := tx.db.Exec("commit")
	if err != nil {
		return errs.ErrCommit
	}

	return nil
}
