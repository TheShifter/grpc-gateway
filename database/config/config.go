package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	driver := "mysql"
	dsn := "root:root@tcp(localhost:3306)/country_schema?charset=utf8"
	db, err = sql.Open(driver, dsn)
	if err != nil {
		panic("InitMysqlDB error" + err.Error())
	}
	return db, nil
}
