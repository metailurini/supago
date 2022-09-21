package util

import (
	"github.com/metailurini/supago/config"
	"github.com/metailurini/supago/database/postgresql"
	"github.com/metailurini/supago/setupcfg"
)

func PostgreSQLLoadEnvConfig(ps postgresql.PostgreSQLSetup, vs config.ViberSetup, configFuncs ...func(setupcfg.Config) error) error {
	configFuncs = append(configFuncs, config.Viper_EnvPostgresql)
	for _, configFunc := range configFuncs {
		if err := vs.Apply(configFunc); err != nil {
			return err
		}
	}

	if err := ps.LoadConfig(vs.GetConfig()); err != nil {
		return err
	}

	return nil
}
