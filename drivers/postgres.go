package drivers

import (
	"database/sql"
	"strings"
)

const (
	del    = ","
	quotes = "'"
)

type postgres struct {
	db *sql.DB
}

func (p *postgres) setConn(db *sql.DB) {
	p.db = db
}

func (p *postgres) save(table string, columns []string, values []interface{}) (interface{}, error) {
	vals := []interface{}{quote(table), quoteAndJoin(columns)}
	vals = append(vals, values...)

	rows, err := p.db.Query("INSERT INTO %s(%s) VALUES (%s) RETURNING ID", vals)
	if err != nil {
		return nil, err
	}

	var id interface{}
	rows.Scan(&id)

	return id, nil
}

// quote escapes a single string
func quote(str string) string {
	return quotes + str + quotes
}

// quote escapes and join a slice of string
func quoteAndJoin(columns []string) string {
	var sql = ""

	if columns != nil {
		for _, col := range columns {
			sql += quote(col) + del
		}

		sql = strings.TrimSuffix(sql, del)
	}

	return sql
}
