package config

import (
	"context"
	"errors"

	"github.com/metailurini/supago/setupcfg"

	"github.com/spf13/viper"
)

type ViperCfg interface {
	setupcfg.Config
}

type viperCfg struct {
	viper *viper.Viper
}

func NewViperCfg(v *viper.Viper) ViperCfg {
	return &viperCfg{viper: v}
}

func (v *viperCfg) Get(key string) interface{} {
	return viper.Get(key)
}

func (v *viperCfg) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (v *viperCfg) Value() interface{} {
	return v.viper
}

type ViberSetup interface {
	setupcfg.Setup
}

type viberSetup struct {
	viper ViperCfg
}

func (v *viberSetup) LoadConfig(cfg setupcfg.Config) error {
	panic("not implemented") // TODO: Implement
}

func (v *viberSetup) GetConfig() setupcfg.Config {
	return v.viper
}

func (v *viberSetup) Apply(setup func(setupcfg.Config) error) error {
	return setup(v.GetConfig())
}

func (v *viberSetup) Value() interface{} {
	return v.viper
}

func NewViperSetup() ViberSetup {
	return &viberSetup{
		viper: NewViperCfg(viper.GetViper()),
	}
}

func Viper_EnvPostgresql(c setupcfg.Config) error {
	v, ok := c.Value().(*viper.Viper)
	if !ok {
		return errors.New("can not apply config for viber")
	}

	if err := v.BindEnv("POSTGRES_URI"); err != nil {
		return err
	}

	uri, _ := v.Get("POSTGRES_URI").(string)
	v.Set("postgresql", uri)

	if v.Get("context") == nil {
		v.Set("context", context.Background())
	}

	return nil
}
