package util

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	"github.com/metailurini/supago/database/postgresql/adapter"
	"github.com/metailurini/supago/mocks/mock_adapter"
	"github.com/metailurini/supago/setupcfg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/spf13/viper"
)

func TestPostgreSQLLoadEnvConfig(t *testing.T) {
	postgres := new(mock_adapter.Postgres)
	adapter.RegisterAdapter("mock", postgres)

	type args struct {
		ps postgresql.PostgreSQLSetup
		vs config.ViberSetup
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			name: "missing postgresql uri",
			args: func(t *testing.T) args {
				vs := config.NewViperSetup()
				ps := postgresql.NewPostgreSQLSetup()
				return args{vs: vs, ps: ps}
			},
			wantErr: true,
		},
		{
			name: "happy case: with default context",
			args: func(t *testing.T) args {
				os.Setenv("POSTGRES_URI", "POSTGRES_URI")

				vs := config.NewViperSetup()
				ps := postgresql.NewPostgreSQLSetup()

				err := vs.Apply(func(c setupcfg.Config) error {
					v, ok := c.Value().(*viper.Viper)
					if !ok {
						return errors.New("can not apply config for viber")
					}
					assert.NotNil(t, v)
					return nil
				})
				assert.NoError(t, err)

				postgres.On("Connect", mock.Anything, "POSTGRES_URI").Once().Return(new(mock_adapter.Postgres), nil)
				postgres.On("Config").Once().Return(new(mock_adapter.Postgres))

				return args{vs: vs, ps: ps}
			},
			wantErr: false,
		},
		{
			name: "happy case: with specific context",
			args: func(t *testing.T) args {
				os.Setenv("POSTGRES_URI", "POSTGRES_URI")

				vs := config.NewViperSetup()
				ps := postgresql.NewPostgreSQLSetup()
				err := vs.Apply(func(c setupcfg.Config) error {
					v, ok := c.Value().(*viper.Viper)
					if !ok {
						return errors.New("can not apply config for viber")
					}
					assert.NotNil(t, v)
					v.Set("context", context.Background())
					return nil
				})
				assert.NoError(t, err)

				postgres.On("Connect", mock.Anything, "POSTGRES_URI").Once().Return(new(mock_adapter.Postgres), nil)
				postgres.On("Config").Once().Return(new(mock_adapter.Postgres))

				return args{vs: vs, ps: ps}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := PostgreSQLLoadEnvConfig(tArgs.ps, tArgs.vs)

			if (err != nil) != tt.wantErr {
				t.Fatalf("PostgreSQLLoadEnvConfig error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
