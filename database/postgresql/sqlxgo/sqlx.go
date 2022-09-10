package sqlxgo

import (
	"context"

	"github.com/metailurini/supago/database/postgresql/adapter"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLX struct {
	Conn *sqlx.DB
}

func init() {
	adapter.RegisterAdapter("sqlx", &SQLX{})
}

func (s *SQLX) Connect(ctx context.Context, uri string) (interface{}, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", uri)
	if err != nil {
		return nil, err
	}
	s.Conn = db
	return s.Conn, nil
}

func (s *SQLX) Config() interface{} {
	return s.Conn
}
