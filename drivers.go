package dkeepr

import (
	"database/sql"

	"bitbucket.org/mathsalmi/dkeepr/drivers"
	"bitbucket.org/mathsalmi/dkeepr/drivers/postgres"
	"bitbucket.org/mathsalmi/dkeepr/errors"
)

// OrmDriver is a DB driver compatible with this ORM
type ormDriver interface {
	DB() *sql.DB
	Name() string

	Begin() (drivers.Transaction, error)
	Save(e drivers.Entity) (interface{}, error)
	Delete(e drivers.Entity) error
}

// NewDriver returns an instance of a given DB driver
func newDriver(name string, db *sql.DB) (driver ormDriver, err error) {
	switch name {
	case "postgres":
		driver = postgres.New(db)
	default:
		err = errs.ErrDriverNotSupported
	}

	return
}
