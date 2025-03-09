package helper

import (
	"strconv"
	"strings"
	"time"
)

func ReplaceQueryParams(namedQuery string, params map[string]any) (string, []any) {
	var (
		i    int = 1
		args []any
	)

	for k, v := range params {
		if k != "" {
			oldsize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
			}
		}
	}

	return namedQuery, args
}

func CheckPostgresTimestamp(value any) bool {
	// Check if the value is a string
	str, ok := value.(string)
	if !ok {
		return false // Not a string
	}

	// Try parsing it using common PostgreSQL timestamp formats
	formats := []string{
		"2006-01-02 15:04:05",    // TIMESTAMP WITHOUT TIME ZONE
		"2006-01-02 15:04:05-07", // TIMESTAMP WITH TIME ZONE
		time.RFC3339,             // ISO 8601 (e.g., "2023-09-18T12:34:56Z")
	}

	for _, format := range formats {
		if _, err := time.Parse(format, str); err == nil {
			return true // Successfully parsed
		}
	}

	return false // No matching format
}
