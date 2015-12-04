package postgres

import "strings"

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
