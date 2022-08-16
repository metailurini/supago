package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"

	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

var yamlCfg = []byte(`postgres: postgresql://postgres:this13dewevwk454f3f25424523f@db.ddhcxdcfqdhkyxrbpfnh.supabase.co:5432/postgres`)

func main() {
	viper.Set("context", context.Background())
	viper.SetConfigType("YAML")
	if err := viper.ReadConfig(bytes.NewBuffer(yamlCfg)); err != nil {
		panic(err)
	}

	ps := postgresql.NewPostgreSQL()
	if err := ps.LoadConfig(config.NewViperCfg()); err != nil {
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
