package postgresql

import (
	"context"
	"errors"

	"github.com/metailurini/supago/setupcfg"

	"github.com/jackc/pgx/v4"
)

var (
	contextErr     = errors.New("can not get context")
	postgresURIErr = errors.New("can not get URI")
)

type PostgreSQLSetup interface {
	setupcfg.Setup
}

type postgresql struct {
	cfg  setupcfg.Config
	conn *pgx.Conn
}

func NewPostgreSQL() PostgreSQLSetup {
	return new(postgresql)
}

func (p *postgresql) LoadConfig(cfg setupcfg.Config) error {
	ctx, ok := cfg.Get("context").(context.Context)
	if !ok {
		return contextErr
	}

	uri, ok := cfg.Get("postgres").(string)
	if !ok {
		return postgresURIErr
	}

	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return err
	}

	p.conn = conn
	p.cfg = &postgreSQLConfig{
		conn: conn.Config(),
	}
	return nil
}

func (p *postgresql) GetConfig() setupcfg.Config {
	return p.cfg
}

func (p *postgresql) Apply(setup func(setupcfg.Config)) {
	setup(p.cfg)
}

func (p *postgresql) CoreValue() interface{} {
	return p.conn
}

type PostgreSQLConfig interface {
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

func (p *postgreSQLConfig) CoreValue() interface{} {
	return p.Config()
}
