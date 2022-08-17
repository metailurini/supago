package config

import (
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

func (v *viberSetup) Apply(setup func(setupcfg.Config)) {
	setup(v.GetConfig())
}

func (v *viberSetup) Value() interface{} {
	return v.viper
}

func NewViperSetup() ViberSetup {
	return &viberSetup{
		viper: NewViperCfg(viper.GetViper()),
	}
}
