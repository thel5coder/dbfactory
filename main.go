package main

import (
	"dbfactory"
	"fmt"
)

func main() {
	psql := dbfactory.NewPsql()
	err := psql.MakeSQLConnection()
	if err != nil {
		panic(err)
	}
	defer psql.CloseConnection()

	fmt.Println(psql.GetPSQLInstance())
}
