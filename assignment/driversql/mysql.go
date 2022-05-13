package driversql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysqlconfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func ConnectToMySQL(conf Mysqlconfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.New("error connecting to database")
	}

	return db, nil
}
