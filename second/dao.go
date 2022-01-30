package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

var errNotFound = errors.New("not found")

func Dao(query string) error {
	err := mockError()

	if err == sql.ErrNoRows {
		return errors.Wrapf(errNotFound, "data not found, sql: %s", query)
	} else if err != nil {
		return errors.Wrapf(err, "db query error, sql: %s", query)
	}
	return nil
}

func mockError() error {
	return sql.ErrNoRows
}

func IsNoRow(err error) bool {
	return errors.Is(err, errNotFound)
}
