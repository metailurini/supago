package main

import (
	"context"
	"fmt"
	"time"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	_ "github.com/metailurini/supago/database/postgresql/sqlxgo"
	"github.com/metailurini/supago/util"

	"github.com/jmoiron/sqlx"
)

// set up server with config is map and setup is http.server
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Minute))
	defer cancel()

	vs := config.NewViperSetup()
	ps := postgresql.NewPostgreSQLSetup()

	vs.GetConfig().Set("context", ctx)

	if err := util.PostgreSQLLoadEnvConfig(ps, vs); err != nil {
		panic(err)
	}

	conn := ps.Value().(*sqlx.DB)
	defer conn.Close()

	count := 0
	stmt := "select count(*) from auth.users"
	err := conn.QueryRow(stmt).Scan(&count)
	if err != nil {
		panic(err)
	}

	fmt.Printf("counted users: %v\n", count)
}
