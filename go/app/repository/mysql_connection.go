package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlConnection(user string, passwd string, addr string, dbName string) *MysqlConnection {
	return &MysqlConnection{
		User:   user,
		Passwd: passwd,
		Addr:   addr,
		DBName: dbName,
	}
}

type MysqlConnection struct {
	User   string
	Passwd string
	Addr   string
	DBName string
}

func (dbc *MysqlConnection) Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 dbc.User,
		Passwd:               dbc.Passwd,
		Net:                  "tcp",
		Addr:                 dbc.Addr,
		DBName:               dbc.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, err
}
