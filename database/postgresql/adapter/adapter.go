package adapter

import (
	"context"
	"sync"
)

var (
	adapterMu sync.RWMutex
	adtName   string
	adt       Postgres
)

type Postgres interface {
	Connect(ctx context.Context, uri string) (interface{}, error)
	Config() interface{}
}

func GetAdapter() (string, Postgres) {
	return adtName, adt
}

func RegisterAdapter(adapterPostgres string, adapter Postgres) {
	adapterMu.Lock()
	defer adapterMu.Unlock()

	adtName = adapterPostgres
	adt = adapter
}
