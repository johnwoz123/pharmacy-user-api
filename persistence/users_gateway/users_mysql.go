package users_gateway

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_DbUsername = "mysql_DbUsername"
	mysql_DbPassword = "mysql_DbPassword"
	mysql_DbHost     = "mysql_DbHost"
	mysql_Schema     = "mysql_Schema"
)

var (
	Client          *sql.DB
	mysqlDbUsername = os.Getenv(mysql_DbUsername)
	mysqlDbPassword = os.Getenv(mysql_DbPassword)
	mysqlDbHost     = os.Getenv(mysql_DbHost)
	mysqlSchema     = os.Getenv(mysql_Schema)
)

func init() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysqlDbUsername,
		mysqlDbPassword,
		mysqlDbHost,
		mysqlSchema)
	var err error

	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal(err)
	}

	if err = Client.Ping(); err != nil {
		log.Println(err)
	}
	log.Println("connected to database ........ ")

}
