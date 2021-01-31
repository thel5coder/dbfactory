package dbfactory

import (
	"database/sql"
	"strconv"
)

func stringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}

type DefaultConfig struct {
	Host                  string
	DbName                string
	User                  string
	Password              string
	Port                  string
	MaxConnection         int
	MaxIdleConnection     int
	MaxLifeTimeConnection int
}

var (
	psqlConn      *sql.DB
	mysqlConn     *sql.DB
	dbTransaction *sql.Tx
)

type Database interface {
	MakeConnection() (err error)

	MakeTransaction() (err error)

	CloseConnection()

	ExecuteRow(sqlStatement string, arguments []interface{}) (row *sql.Row, err error)

	ExecuteRows(sqlStatement string, argument []interface{}) (rows *sql.Rows, err error)

	GetPSQLInstance() (db *sql.DB)
}
