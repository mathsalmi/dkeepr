package postgres

import (
	"database/sql"
	"fmt"
	"strings"
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
func (p *Postgres) Save(table string, columns []string, values []interface{}) (interface{}, error) {

	placeholders := makePlaceholders(len(columns))

	sql := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s) RETURNING "ID"`, quote(table), quoteAndJoin(columns), strings.Join(placeholders, ","))

	rows, err := p.db.Query(sql, values...)
	if err != nil {
		return nil, err
	}

	var id interface{}
	rows.Scan(&id)

	return id, nil
}

// MakePlaceholders returns a slice of placeholders
func makePlaceholders(total int) []string {
	var out = make([]string, total)
	for i := 0; i < total; i++ {
		out[i] = "$" + fmt.Sprintf("%d", i+1)
	}

	return out
}
