package config

type Setup interface {
	LoadConfig(cfg Config) error
	Apply(setup func(Config) Config)
	CoreValue() interface{}
}

type Config interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}
