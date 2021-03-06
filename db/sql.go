package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/somewhere/config"
)

var SqlDb *sql.DB

func InitSQLDatabase() error {

	dsn := getDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("123", err)
		return err
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("123", err)
		return err
	}
	SqlDb = db
	return nil
}

func getDSN() string {
	// DSN format user:password@tcp(your-amazonaws-uri.com:3306)/dbname
	ret := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.DbConfig.SQLName, config.Config.DbConfig.SQLPassword,
		config.Config.DbConfig.SQLAddress, config.Config.DbConfig.SQLDbName)
	fmt.Println("ret", ret)
	return ret
}
