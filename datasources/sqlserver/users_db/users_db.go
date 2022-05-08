package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
)

const (
	mssql_users_username = "mssql_users_username"
	mssql_users_password = "mssql_users_password"
	mssql_users_server   = "mssql_users_server"
	mssql_users_port     = "mssql_users_port"
	mssql_users_database = "mssql_users_database"
)

var (
	Client   *sql.DB
	password = os.Getenv(mssql_users_password)
	user     = os.Getenv(mssql_users_username)
	port     = os.Getenv(mssql_users_port)
	server   = os.Getenv(mssql_users_server)
	database = os.Getenv(mssql_users_database)
)

func init() {
	datasourceName := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)
	var err error
	Client, err = sql.Open("mssql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connection successful")
}
