package mssql_utils

import (
	"fmt"
	"github.com/bm1905/bookstore_users_api/utils/errors_utils"
	"github.com/denisenkom/go-mssqldb"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors_utils.RestError {
	sqlErr, ok := err.(mssql.Error)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors_utils.NewNotFoundError("no record matches given id")
		}
		return errors_utils.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.SQLErrorNumber() {
	case 2627:
		return errors_utils.NewBadRequestError(fmt.Sprintf("already taken"))
	}
	return errors_utils.NewInternalServerError("error processing request")
}
