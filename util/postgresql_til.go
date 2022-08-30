package util

import (
	"context"
	"errors"
	"fmt"

	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	"github.com/metailurini/supago/setupcfg"
	"github.com/spf13/viper"
)

func PostgreSQLLoadEnvConfig(ps postgresql.PostgreSQLSetup, vs config.ViberSetup) error {
	if err := vs.Apply(func(c setupcfg.Config) error {
		v, ok := c.Value().(*viper.Viper)
		if !ok {
			return errors.New("can not apply config for viber")
		}

		if err := v.BindEnv("POSTGRES_URI"); err != nil {
			return err
		}

		uri, ok := v.Get("POSTGRES_URI").(string)
		if !ok {
			return errors.New("can not get postgres uri")
		}

		v.Set("postgresql", uri)
		v.Set("context", context.Background())

		return nil
	}); err != nil {
		return err
	}

	if err := ps.LoadConfig(vs.GetConfig()); err != nil {
		return err
	}
	fmt.Printf("ps.Value(): %v\n", ps.Value())

	return nil
}
