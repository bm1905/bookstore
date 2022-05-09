package users

import (
	"fmt"
	"github.com/bm1905/bookstore_users_api/datasources/sqlserver/users_db"
	"github.com/bm1905/bookstore_users_api/logger"
	"github.com/bm1905/bookstore_users_api/utils/errors_utils"
	"github.com/bm1905/bookstore_users_api/utils/mssql_utils"
)

const (
	queryInsertUser      = "INSERT INTO users(first_name, last_name, email, date_created, password, status) OUTPUT INSERTED.ID VALUES(?, ?, ?, ?, ?, ?);"
	queryUpdateUser      = "UPDATE users SET first_name=?, last_name=? WHERE id=?;"
	queryGetUser         = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryGetAllUsers     = "SELECT id, first_name, last_name, email, date_created, status FROM users;"
	queryDeleteUser      = "DELETE FROM users WHERE id=?;"
	queryGetUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while preparing statement", err)
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("database error", err)
		return mssql_utils.ParseError(err)
	}

	return nil
}

func (user *User) GetAll() ([]User, *errors_utils.RestError) {
	stmt, err := users_db.Client.Prepare(queryGetAllUsers)
	if err != nil {
		logger.Error("error while preparing statement", err)
		return nil, errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("database error", err)
		return nil, errors_utils.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("database error", err)
			return nil, mssql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors_utils.NewNotFoundError(fmt.Sprintf("no users found"))
	}

	return results, nil
}

func (user *User) Save() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error while preparing statement", err)
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var lastInsertId int64
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status).Scan(&lastInsertId)
	if err != nil {
		logger.Error("database error", err)
		return mssql_utils.ParseError(err)
	}

	user.Id = lastInsertId

	return nil
}

func (user *User) Update() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error while preparing statement", err)
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Id)
	if err != nil {
		logger.Error("database error", err)
		return mssql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors_utils.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error while preparing statement", err)
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("database error", err)
		return mssql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors_utils.RestError) {
	stmt, err := users_db.Client.Prepare(queryGetUserByStatus)
	if err != nil {
		logger.Error("error while preparing statement", err)
		return nil, errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("database error", err)
		return nil, errors_utils.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("database error", err)
			return nil, mssql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors_utils.NewNotFoundError(fmt.Sprintf("no users found for status %s", status))
	}

	return results, nil
}
