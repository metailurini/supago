package pgxgo

import (
	"context"

	"github.com/metailurini/supago/database/postgresql/adapter"

	"github.com/jackc/pgx/v4"
)

type PGX struct {
	Conn *pgx.Conn
}

func init() {
	adapter.RegisterAdapter("pgx", &PGX{})
}

func (s *PGX) Connect(ctx context.Context, uri string) (interface{}, error) {
	db, err := pgx.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}
	s.Conn = db
	return s.Conn, nil
}

func (s *PGX) Config() interface{} {
	return s.Conn
}
