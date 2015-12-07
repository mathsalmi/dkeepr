package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"bitbucket.org/mathsalmi/dkeepr/drivers"
	"bitbucket.org/mathsalmi/dkeepr/errors"
)

const (
	del    = ","
	quotes = `"`
)

// Postgres is the driver for the Postgresql DB
type Postgres struct {
	db *sql.DB
}

// New returns an instance of Postgres
func New(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

// DB returns an instance of the connection with the DB
func (p *Postgres) DB() *sql.DB {
	return p.db
}

// Name returns the driver name
func (p *Postgres) Name() string {
	return "postgres"
}

// Save saves an entity
func (p *Postgres) Save(e drivers.Entity) (interface{}, error) {

	var sql = ""

	if IsPkEmpty(e) {

		placeholders := makePlaceholders(len(e.Columns()))

		sql = fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s) RETURNING %s`,
			quote(e.Tablename()),
			quoteAndJoin(e.Columns()),
			strings.Join(placeholders, del),
			quoteAndJoin(e.Pk()))
	} else {
		fields := makePairPlaceholders(e.Columns())
		ids := makePairPlaceholders(e.Pk())
		sql = fmt.Sprintf(`UPDATE %s SET %s WHERE %s RETURNING %s`,
			quote(e.Tablename()),
			strings.Join(fields, del),
			strings.Join(ids, " AND "),
			strings.Join(ids, " AND "))
	}

	fmt.Println(sql)

	row := p.db.QueryRow(sql, e.Values()...)
	if row == nil {
		return nil, errs.ErrNoResult
	}

	var id interface{}
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return id, nil
}

// Delete deletes an entity
func (p *Postgres) Delete(e drivers.Entity) error {
	// check ids
	if IsPkEmpty(e) {
		return errs.ErrNoPk
	}

	placeholders := makePairPlaceholders(e.Pk())

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", quote(e.Tablename()), strings.Join(placeholders, " AND "))

	_, err := p.db.Exec(sql, e.Pkvalues()...)
	if err != nil {
		return err
	}

	return nil
}

// IsPkEmpty returns false if at least one element of ID slice
// is different than the zero value for its type
//
// eg.: zero value for int is 0, it checks if id != 0
//      zero value for string is "", it checks if id != ""
//      and so on…
func IsPkEmpty(e drivers.Entity) bool {
	values := e.Pkvalues()
	for _, val := range values {
		zero := reflect.Zero(reflect.TypeOf(val))
		if val != zero.Interface() {
			return false
		}
	}

	return true
}
