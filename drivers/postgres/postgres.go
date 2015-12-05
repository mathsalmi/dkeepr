package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"bitbucket.org/mathsalmi/dkeepr/drivers"
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

	placeholders := makePlaceholders(len(e.Columns()))

	sql := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s) RETURNING %s`,
		quote(e.Tablename()),
		quoteAndJoin(e.Columns()),
		strings.Join(placeholders, del),
		quoteAndJoin(e.Pk()))

	row := p.db.QueryRow(sql, e.Values()...)
	if row == nil {
		return nil, drivers.ErrNoResult
	}

	var id interface{}
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return id, nil
}

// Delete deletes an entity
func (p *Postgres) Delete(e drivers.Entity) error {
	placeholders := makePairPlaceholders(e.Pk())

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", quote(e.Tablename()), strings.Join(placeholders, " AND "))

	_, err := p.db.Exec(sql, e.Pkvalues()...)
	if err != nil {
		return err
	}

	return nil
}
