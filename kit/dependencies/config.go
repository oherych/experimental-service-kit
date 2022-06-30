package dependencies

type Config interface {
	GetBaseConfig() BaseConfig
}

type BaseConfig struct {
	AppName  string
	Debug    bool
	HttpPort string `default:":8000"`
	GRPCPort string `default:":50051"`
	HTTPPort string `default:":8080"`
}
