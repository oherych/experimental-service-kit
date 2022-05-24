package kit

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
)

type Builder[Dep dependencies.Locator] func(dep Dep, sr grpc.ServiceRegistrar) error

type GRPC[Dep dependencies.Locator] struct {
	Builder func(dep Dep, sr grpc.ServiceRegistrar) error
}

func (m GRPC[Dep]) Server(ctx context.Context, log zerolog.Logger, dep Dep, bc dependencies.BaseConfig) error {
	log.Info().Str("port", bc.GRPCPort).Msg("[SYS] Starting gRPC server on port")

	lis, err := net.Listen("tcp", bc.GRPCPort)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	go func() {
		<-ctx.Done()

		log.Info().Msg("[SYS] Stop gRPC server")

		s.GracefulStop()
	}()

	if err := m.Builder(dep, s); err != nil {
		return err
	}

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
