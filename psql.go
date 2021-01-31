package dbfactory

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type Psql struct {
	DefaultConfig
	SslMode     string
	SslCert     string
	SslKey      string
	SslRootCert string
}

//creating new instance of postgresql connection
func NewPsql() Database {
	//set default maximum open connection
	maxCon := stringToInt(os.Getenv("DB_MAX_CONNECTION"))
	if maxCon == 0 {
		maxCon = 3
	}

	//set default maximum idle connection
	maxIdleConn := stringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION"))
	if maxIdleConn == 0 {
		maxIdleConn = 3
	}

	//set default maximum life time connection
	maxLifeTimeConn := stringToInt(os.Getenv("DB_MAX_LIFE_CONNECTION"))
	if maxLifeTimeConn == 0 {
		maxLifeTimeConn = 3
	}

	return &Psql{
		DefaultConfig: DefaultConfig{
			Host:                  os.Getenv("DB_HOST"),
			DbName:                os.Getenv("DB_NAME"),
			User:                  os.Getenv("DB_USER_NAME"),
			Password:              os.Getenv("DB_PASSWORD"),
			Port:                  os.Getenv("DB_PORT"),
			MaxConnection:         maxCon,
			MaxIdleConnection:     maxIdleConn,
			MaxLifeTimeConnection: maxLifeTimeConn,
		},
		SslMode:     os.Getenv("DB_SSL_MODE"),
		SslCert:     os.Getenv("DB_SSL_CERT"),
		SslKey:      os.Getenv("DB_SSL_KEY"),
		SslRootCert: os.Getenv("DB_SSL_ROOT_CERT"),
	}
}

func (psql *Psql) MakeConnection() (err error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=UTC", psql.DefaultConfig.User, psql.DefaultConfig.Password,
		psql.DefaultConfig.Host, psql.DefaultConfig.Port, psql.DefaultConfig.DbName, psql.SslMode,
	)

	if psql.SslMode == "require" {
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=UTC&sslcert=%s&sslkey=%s&sslrootcert=%s",
			psql.DefaultConfig.User, psql.DefaultConfig.Password,
			psql.DefaultConfig.Host, psql.DefaultConfig.Port, psql.DefaultConfig.DbName, psql.SslMode,
			psql.SslCert, psql.SslKey, psql.SslRootCert,
		)
	}

	psqlConn, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer psqlConn.Close()
	psqlConn.SetMaxOpenConns(psql.MaxConnection)
	psqlConn.SetMaxIdleConns(psql.MaxIdleConnection)
	psqlConn.SetConnMaxLifetime(time.Duration(psql.MaxLifeTimeConnection) * time.Second)

	return nil
}

func (psql *Psql) CloseConnection() {
	psqlConn.Close()
}

func (psql *Psql) MakeTransaction() (err error) {
	if psqlConn == nil {
		return errors.New("connection is null")
	}

	dbTransaction, err := psqlConn.Begin()
	if err != nil {
		dbTransaction.Rollback()
		return err
	}

	return nil
}

func (psql *Psql) ExecuteRow(sqlStatement string, arguments []interface{}) (row *sql.Row, err error) {
	panic("implement me")
}

func (psql *Psql) ExecuteRows(sqlStatement string, argument []interface{}) (rows *sql.Rows, err error) {
	panic("implement me")
}

func (psql *Psql) GetPSQLInstance() (db *sql.DB) {
	return psqlConn
}
