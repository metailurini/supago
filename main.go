package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	_ "github.com/metailurini/supago/database/postgresql/sqlxgo"
	"github.com/metailurini/supago/setupcfg"
	"github.com/metailurini/supago/util"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Minute))
	defer cancel()

	vs := config.NewViperSetup()
	ps := postgresql.NewPostgreSQLSetup()

	if err := vs.Apply(func(c setupcfg.Config) error {
		v, ok := c.Value().(*viper.Viper)
		if !ok {
			return errors.New("can not apply config for viber")
		}

		v.Set("context", ctx)

		return nil
	}); err != nil {
		panic(err)
	}

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
