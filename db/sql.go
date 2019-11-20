package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.bj.sensetime.com/SenseGo/camera-kit/config"
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
	ret := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.DbConfig.Name, config.Config.DbConfig.Password,
		config.Config.DbConfig.Address, config.Config.DbConfig.DbName)
	fmt.Println("ret", ret)
	return ret
}
