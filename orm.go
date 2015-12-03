package main

import (
	"database/sql"

	"bitbucket.org/mathsalmi/dkeepr/drivers"
)

// Dkeepr is an ORM object
type Dkeepr struct {
	conn   *sql.DB
	driver string
	url    string
}

// NewDkeepr returns a new Dkeepr
func NewDkeepr(driver, url string) (*Dkeepr, error) {
	orm := &Dkeepr{}

	err := orm.SetDriver(driver)
	if err != nil {
		return nil, err
	}

	orm.url = url

	return orm, nil
}

// Driver tells the DB driver currently in use
func (d *Dkeepr) Driver() string {
	return d.driver
}

// SetDriver tells the ORM which DB driver to use
func (d *Dkeepr) SetDriver(driver string) error {
	if !drivers.IsDriverSupported(driver) {
		return ErrDriverNotSupported
	}

	d.driver = driver

	return nil
}

// Open opens the connection with the DB
func (d *Dkeepr) Open() error {
	db, err := sql.Open(d.driver, d.url)
	if err != nil {
		return err
	}

	d.conn = db

	return nil
}

// Close closes the open connection
func (d *Dkeepr) Close() error {
	return d.conn.Close()
}
