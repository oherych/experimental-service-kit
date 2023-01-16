package grpc_listener

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
)

type Builder[Dep dependencies.Locator] func(dep Dep, sr grpc.ServiceRegistrar) error

type GRPC[Conf dependencies.Config, Dep dependencies.Locator] struct {
	Swagger      []byte
	InitConfig   func(conf Conf) Config
	Builder      func(sr grpc.ServiceRegistrar, dep Dep) error
	HTTPHandlers []func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
}

func (m GRPC[Conf, Dep]) Server(ctx context.Context, log *zerolog.Logger, dep Dep, global Conf) error {
	config := m.InitConfig(global)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return m.buildGRPC(ctx, log, dep, config)
	})

	return g.Wait()
}

func (m GRPC[Conf, Dep]) buildGRPC(ctx context.Context, log *zerolog.Logger, dep Dep, bc Config) error {
	log.Info().Str("port", bc.GRPCPort).Msg("[SYS] Starting gRPC cmd")

	s := m.newServer(ctx, log)

	if err := m.Builder(s, dep); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", bc.GRPCPort)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}

func (m GRPC[Conf, Dep]) newServer(ctx context.Context, log *zerolog.Logger) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	go func() {
		<-ctx.Done()

		log.Info().Msg("[SYS] Stop gRPC cmd")

		s.GracefulStop()
	}()

	return s
}
