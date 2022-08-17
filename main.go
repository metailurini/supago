package main

import (
	"context"
	"fmt"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	"github.com/metailurini/supago/util"

	"github.com/jackc/pgx/v4"
)

func main() {
	vs := config.NewViperSetup()
	ps := postgresql.NewPostgreSQLSetup()

	if err := util.PostgreSQLLoadEnvConfig(ps, vs); err != nil {
		panic(err)
	}

	conn := ps.Value().(*pgx.Conn)
	defer conn.Close(context.Background())

	var userid int
	var query string
	stmt := "select userid, query from pg_stat_statements limit 1"
	err := conn.QueryRow(context.Background(), stmt).Scan(&userid, &query)
	if err != nil {
		panic(err)
	}

	fmt.Printf("userid: %v\n", userid)
	fmt.Printf("query: %v\n", query)
}
