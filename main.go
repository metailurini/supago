package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

// load from env
// override config from env
// return core package

var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
postgres: postgresql://postgres:this13dewevwk454f3f25424523f@db.ddhcxdcfqdhkyxrbpfnh.supabase.co:5432/postgres
`)

type SetupWrapper interface {
	LoadConfig(cfg Config) error
	Apply(setup func(Config) Config)
	CoreValue() interface{}
}

type Config interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

type Postgre struct {
	conn *pgx.Conn
}

func (p *Postgre) LoadConfig(cfg Config) error {
	ctx, ok := cfg.Get("context").(context.Context)
	if !ok {
		return errors.New("can not get context")
	}

	uri, ok := cfg.Get("postgres").(string)
	if !ok {
		return errors.New("can not get uri")
	}

	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return err
	}

	p.conn = conn
	return nil
}

func (p *Postgre) Apply(setup func(Config) Config) {
	panic("not implemented") // TODO: Implement
}

func (p *Postgre) CoreValue() interface{} {
	return p.conn
}

type ViperWrapper struct{}

func (v *ViperWrapper) Get(key string) interface{} {
	return viper.Get(key)
}

func (v *ViperWrapper) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func main() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	viper.Set("context", context.Background())

	// any approach to require this configuration into your program.

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		panic(err)
	}

	p := Postgre{}
	if err := p.LoadConfig(&ViperWrapper{}); err != nil {
		panic(err)
	}

	fmt.Printf("viper.GetString(\"postgres\"): %v\n", viper.GetString("postgres"))

	conn := p.CoreValue().(*pgx.Conn)
	// conn, err := pgx.Connect(context.Background(), viper.GetString("postgres"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var userid int
	var query string
	err = conn.QueryRow(context.Background(), "select userid, query from pg_stat_statements limit 1").Scan(&userid, &query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(userid, query)
}
