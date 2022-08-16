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

var yamlConfig = []byte(`postgres: postgresql://postgres:this13dewevwk454f3f25424523f@db.ddhcxdcfqdhkyxrbpfnh.supabase.co:5432/postgres`)

func main() {
	viper.SetConfigType("YAML")
	viper.Set("context", context.Background())

	if err := viper.ReadConfig(bytes.NewBuffer(yamlConfig)); err != nil {
		panic(err)
	}

	ps := postgresql.NewPostgreSQL()
	showPostgre(ps)
}

type ViperWrapper struct{}

func (v *ViperWrapper) Get(key string) interface{} {
	return viper.Get(key)
}

func (v *ViperWrapper) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func showPostgre(su config.Setup) {
	if err := su.LoadConfig(&ViperWrapper{}); err != nil {
		panic(err)
	}

	conn := su.CoreValue().(*pgx.Conn)
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
