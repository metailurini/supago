package postgresql

import (
	"context"
	"testing"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql/adapter"
	_ "github.com/metailurini/supago/database/postgresql/sqlxgo"
	"github.com/metailurini/supago/mocks/mock_adapter"
	"github.com/metailurini/supago/setupcfg"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostgresql_LoadConfig(t *testing.T) {
	postgres := new(mock_adapter.Postgres)
	adapter.RegisterAdapter("mock", postgres)

	type args struct {
		cfg setupcfg.Config
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *postgresql
		inspect func(r *postgresql, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			name: "context was not set",
			init: func(t *testing.T) *postgresql {
				return NewPostgreSQLSetup().(*postgresql)
			},
			inspect: func(r *postgresql, t *testing.T) {},
			args: func(t *testing.T) args {
				vs := config.NewViperSetup()
				return args{cfg: vs.GetConfig()}
			},
			wantErr: true,
		},
		{
			name: "postgresql was not set",
			init: func(t *testing.T) *postgresql {
				return NewPostgreSQLSetup().(*postgresql)
			},
			inspect: func(r *postgresql, t *testing.T) {},
			args: func(t *testing.T) args {
				vs := config.NewViperSetup()
				err := vs.Apply(func(c setupcfg.Config) error {
					v, _ := c.Value().(*viper.Viper)
					v.Set("context", context.Background())
					return nil
				})
				assert.NoError(t, err)
				return args{cfg: vs.GetConfig()}
			},

			wantErr: true,
		},
		{
			name: "happy case",
			init: func(t *testing.T) *postgresql {
				postgres.On("Connect", mock.Anything, "postgresql").Once().Return(new(mock_adapter.Postgres), nil)
				postgres.On("Config").Once().Return(new(mock_adapter.Postgres))
				return NewPostgreSQLSetup().(*postgresql)
			},
			inspect: func(r *postgresql, t *testing.T) {
				p, ok := r.Value().(*mock_adapter.Postgres)
				assert.NotNil(t, p)
				assert.True(t, ok)
			},
			args: func(t *testing.T) args {
				vs := config.NewViperSetup()
				err := vs.Apply(func(c setupcfg.Config) error {
					v, _ := c.Value().(*viper.Viper)
					v.Set("context", context.Background())
					v.Set("postgresql", "postgresql")
					return nil
				})
				assert.NoError(t, err)
				return args{cfg: vs.GetConfig()}
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			err := receiver.LoadConfig(tArgs.cfg)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("postgresql.LoadConfig error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestPostgresql_GetConfig(t *testing.T) {
	postgres := new(mock_adapter.Postgres)
	adapter.RegisterAdapter("mock", postgres)

	tests := []struct {
		name    string
		init    func(t *testing.T) *postgresql
		inspect func(r *postgresql, t *testing.T) //inspects receiver after test run

		inspectRes func(t *testing.T, cfg setupcfg.Config)
	}{
		{
			name: "happy case: no config",
			init: func(t *testing.T) *postgresql {
				ps := NewPostgreSQLSetup()
				return ps.(*postgresql)
			},
			inspectRes: func(t *testing.T, cfg setupcfg.Config) {
				assert.Nil(t, cfg)
			},
		},
		{
			name: "happy case: with viper config",
			init: func(t *testing.T) *postgresql {

				ps := NewPostgreSQLSetup()
				vs := config.NewViperSetup()
				err := vs.Apply(func(c setupcfg.Config) error {
					v, _ := c.Value().(*viper.Viper)
					v.Set("context", context.Background())
					v.Set("postgresql", "postgresql")
					return nil
				})
				assert.NoError(t, err)

				postgres.On("Connect", mock.Anything, "postgresql").Once().Return(nil, nil)
				postgres.On("Config").Once().Return(new(mock_adapter.Postgres))

				err = ps.LoadConfig(vs.GetConfig())
				assert.NoError(t, err)
				return ps.(*postgresql)
			},
			inspectRes: func(t *testing.T, cfg setupcfg.Config) {
				_, ok := cfg.Value().(*mock_adapter.Postgres)
				assert.True(t, ok)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.GetConfig()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if tt.inspectRes != nil {
				tt.inspectRes(t, got1)
			}
		})
	}
}
