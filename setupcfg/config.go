package setupcfg

type Setup interface {
	LoadConfig(cfg Config) error
	GetConfig() Config
	Apply(setup func(Config))
	CoreValue() interface{}
}

type Config interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	CoreValue() interface{}
}
