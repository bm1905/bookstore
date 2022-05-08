package users

import (
	"github.com/bm1905/bookstore_users_api/datasources/sqlserver/users_db"
	"github.com/bm1905/bookstore_users_api/utils/dates_utils"
	"github.com/bm1905/bookstore_users_api/utils/errors_utils"
	"github.com/bm1905/bookstore_users_api/utils/mssql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) OUTPUT INSERTED.ID VALUES(?, ?, ?, ?);"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=? WHERE id=?;"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

const (
	errorNoRows = "no rows in result set"
)

func (user *User) Get() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mssql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dates_utils.GetNowString()

	var lastInsertId int64
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&lastInsertId)
	if err != nil {
		return mssql_utils.ParseError(err)
	}

	user.Id = lastInsertId

	return nil
}

func (user *User) Update() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Id)
	if err != nil {
		return mssql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mssql_utils.ParseError(err)
	}

	return nil
}
