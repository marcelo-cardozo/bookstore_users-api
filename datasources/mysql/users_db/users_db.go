package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	users_db_username = "users_db_username"
	users_db_password = "users_db_password"
	users_db_host = "users_db_host"
	users_db_dbname = "users_db_dbname"
)

var (
	dbUsername = os.Getenv(users_db_username)
	dbPassword = os.Getenv(users_db_password)
	dbHost = os.Getenv(users_db_host)
	dbName = os.Getenv(users_db_dbname)

	Client *sql.DB
)


func init(){
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, dbHost, dbName)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connected")
}
