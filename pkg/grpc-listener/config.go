package grpc_listener

type Config struct {
	HTTPPort string `default:":8000"`
	GRPCPort string `default:":50051"`
}
