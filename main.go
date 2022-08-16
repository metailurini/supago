package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	"github.com/metailurini/supago/setupcfg"

	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

var yamlCfg = []byte(`postgres: postgresql://postgres:this13dewevwk454f3f25424523f@db.ddhcxdcfqdhkyxrbpfnh.supabase.co:5432/postgres`)

func main() {

	vs := config.NewViperSetup()
	vs.Apply(func(c setupcfg.Config) {
		v := c.CoreValue().(*viper.Viper)

		v.Set("context", context.Background())
		v.SetConfigType("YAML")

		if err := v.ReadConfig(bytes.NewBuffer(yamlCfg)); err != nil {
			panic(err)
		}
	})

	ps := postgresql.NewPostgreSQL()
	if err := ps.LoadConfig(vs.GetConfig()); err != nil {
		panic(err)
	}

	conn := ps.CoreValue().(*pgx.Conn)
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
