package postgres

import (
	"fmt"
	"strings"
)

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

// MakePlaceholders returns a slice of placeholders
func makePlaceholders(total int) []string {
	var out = make([]string, total)
	for i := 0; i < total; i++ {
		out[i] = "$" + fmt.Sprintf("%d", i+1)
	}

	return out
}

// MakePairPlaceholders returns a slice with pair-based placeholder.
// Eg.: something like KEY=$1...
func makePairPlaceholders(keys []string) []string {
	out := make([]string, len(keys))
	for i, key := range keys {
		out[i] = fmt.Sprintf("%s=$%d", quote(key), i+1)
	}

	return out
}
