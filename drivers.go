package dkeepr

import (
	"database/sql"

	"bitbucket.org/mathsalmi/dkeepr/drivers/postgres"
)

// OrmDriver is a DB driver compatible with this ORM
type ormDriver interface {
	DB() *sql.DB
	Name() string

	Save(table string, columns []string, values []interface{}) (interface{}, error)
}

// NewDriver returns an instance of a given DB driver
func newDriver(name string, db *sql.DB) (driver ormDriver, err error) {
	switch name {
	case "postgres":
		driver = postgres.New(db)
	default:
		err = ErrDriverNotSupported
	}

	return
}