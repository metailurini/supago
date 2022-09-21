package main

import (
	"context"
	"fmt"

	"example/sqlc"

	"github.com/jmoiron/sqlx"
	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	_ "github.com/metailurini/supago/database/postgresql/sqlxgo"
	"github.com/metailurini/supago/util"
)

func main() {
	vc := config.NewViperSetup()
	ps := postgresql.NewPostgreSQLSetup()
	ctx := context.Background()

	vc.GetConfig().Set("context", ctx)
	err := util.PostgreSQLLoadEnvConfig(ps, vc)
	if err != nil {
		panic(err)
	}

	db := ps.Value().(*sqlx.DB)
	q := sqlc.New(db)
	i, _ := q.GetAmountTips(ctx)
	fmt.Printf("i: %v\n", i)
}
