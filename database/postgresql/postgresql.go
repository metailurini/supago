package postgresql

import (
	"context"
	"errors"

	"github.com/metailurini/supago/config"

	"github.com/jackc/pgx/v4"
)

type PostgreSQLSetup interface {
	config.Setup
}

type postgresql struct {
	cfg PostgreSQLConfig
	con *pgx.Conn
}

func NewPostgreSQL() PostgreSQLSetup {
	return new(postgresql)
}

func (p *postgresql) LoadConfig(cfg config.Config) error {
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

	p.con = conn
	p.cfg = &postgreSQLConfig{
		conn: conn.Config(),
	}
	return nil
}

func (p *postgresql) Apply(setup func(config.Config) config.Config) {
	setup(p.cfg)
}

func (p *postgresql) CoreValue() interface{} {
	return p.con
}

type PostgreSQLConfig interface {
	config.Config
	Config() *pgx.ConnConfig
}

type postgreSQLConfig struct {
	conn *pgx.ConnConfig
}

func (p *postgreSQLConfig) Get(key string) interface{} {
	panic("not implemented") // TODO: Implement
}

func (p *postgreSQLConfig) Set(key string, value interface{}) {
	panic("not implemented") // TODO: Implement
}

func (p *postgreSQLConfig) Config() *pgx.ConnConfig {
	return p.Config()
}
