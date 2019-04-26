package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	engine   = "mysql"
	username = "root"
	password = "root"
	port     = "3306"
	name     = "country_schema"
	hostName = "localhost"
	unicode  = "utf8"
)

func GetConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, hostName, port, name, unicode)
	db, err := sql.Open(engine, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
