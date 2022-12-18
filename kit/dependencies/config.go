package dependencies

type Config interface {
	GetBaseConfig() Base
	Validate() error
}

type ConfigSetter interface {
	SetBaseConfig(in Base)
}

type Base struct {
	AppName  string `ignored:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info" desc:"Log level"`
}

func (bc Base) GetBaseConfig() Base {
	return bc
}

func (bc *Base) SetBaseConfig(in Base) {
	*bc = in
}
