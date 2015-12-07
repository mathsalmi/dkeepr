package dkeepr

import (
	"database/sql"

	"bitbucket.org/mathsalmi/dkeepr/errors"
)

// Dkeepr is an ORM object
type Dkeepr struct {
	driver ormDriver
}

// NewDkeepr returns a new Dkeepr
func NewDkeepr(driver, url string) (*Dkeepr, error) {
	// open db connection
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	// get driver instance
	ormdriver, err := newDriver(driver, db)
	if err != nil {
		return nil, err
	}

	orm := &Dkeepr{driver: ormdriver}
	return orm, nil
}

// DriverName tells the DB driver currently in use
func (d *Dkeepr) DriverName() string {
	name := ""
	if d.driver != nil {
		name = d.driver.Name()
	}
	return name
}

// Close closes the open connection
func (d *Dkeepr) Close() error {
	switch {
	case d.driver == nil:
		return errs.ErrDriverNotChosen
	case d.driver.DB() == nil:
		return errs.ErrConnNotOpen
	}

	return d.driver.DB().Close()
}
