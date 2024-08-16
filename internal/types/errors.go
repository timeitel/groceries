package types

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrSQLUnique   = errors.New("UNIQUE constraint failed")
	ErrSQLNotFound = errors.New("Not found")
)

var errMsgSQLUnique = "error code = 1: Error fetching next row: SQLite failure: `UNIQUE constraint failed:"

func NewSQLError(context string, err error) error {
	wrappedErr := fmt.Errorf("Original error: %w", err)

	isUniqueErr := strings.Contains(err.Error(), errMsgSQLUnique)
	if isUniqueErr {
		return fmt.Errorf("Context error: %s\nSQL error: %w\n%w\n", context, ErrSQLUnique, wrappedErr)
	}

	return fmt.Errorf("Context error: %s\n%w\n", context, wrappedErr)
}
