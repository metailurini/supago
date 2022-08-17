package setupcfg

type Setup interface {
	LoadConfig(cfg Config) error
	GetConfig() Config
	Apply(setup func(Config) error) error
	Value() interface{}
}

type Config interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Value() interface{}
}
