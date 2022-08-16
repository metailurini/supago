package config

import (
	"github.com/metailurini/supago/setupcfg"

	"github.com/spf13/viper"
)

type ViperCfg interface {
	setupcfg.Config
}

type viperCfg struct{}

func NewViperCfg() ViperCfg {
	return new(viperCfg)
}

func (v *viperCfg) Get(key string) interface{} {
	return viper.Get(key)
}

func (v *viperCfg) Set(key string, value interface{}) {
	viper.Set(key, value)
}
