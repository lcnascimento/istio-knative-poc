package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"

	"google.golang.org/grpc"
	ggrpc "google.golang.org/grpc"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"
)

// ServerInput ...
type ServerInput struct {
	Port        int
	Tracer      trace.Tracer
	Logger      infra.LogProvider
	Registrator func(srv ggrpc.ServiceRegistrar)
}

// Server ...
type Server struct {
	in ServerInput
}

// NewServer ...
func NewServer(in ServerInput) (*Server, error) {
	if in.Port <= 80 {
		return nil, errors.New("Port value should be greater than or equal to 80")
	}

	if in.Tracer == nil {
		return nil, errors.New("Missing required dependency: Tracer")
	}

	if in.Logger == nil {
		return nil, errors.New("Missing required dependency: Logger")
	}

	return &Server{in: in}, nil
}

// Listen ...
func (s Server) Listen(ctx context.Context) error {
	ctx, span := s.in.Tracer.Start(ctx, "infra.grpc.Listen")
	defer span.End()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.in.Port))
	if err != nil {
		s.in.Logger.Critical(ctx, errors.New("can not initialize gRPC server %s", err.Error()))
		return err
	}

	srv := ggrpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.ChainStreamInterceptor(
			otelgrpc.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		),
	)

	s.in.Registrator(srv)
	grpc_prometheus.Register(srv)

	s.in.Logger.Debug(ctx, "gRPC server started")
	if err := srv.Serve(lis); err != nil {
		s.in.Logger.Critical(ctx, errors.New("could not initialize grpc server: %s", err.Error()))
		return err
	}

	return nil
}

// ClientInput ...
type ClientInput struct {
	ServerAddress string

	Tracer trace.Tracer
	Logger infra.LogProvider
}

// Client ...
type Client struct {
	in ClientInput
}

// NewClient ...
func NewClient(in ClientInput) (*Client, error) {
	return &Client{in: in}, nil
}

// Connect ...
func (c Client) Connect(ctx context.Context) (*grpc.ClientConn, error) {
	ctx, span := c.in.Tracer.Start(ctx, "domain.segments.Connect")
	defer span.End()

	c.in.Logger.Info(ctx, "Connecting to gRPc server")
	conn, err := grpc.Dial(
		c.in.ServerAddress,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(),
			grpc_prometheus.UnaryClientInterceptor,
		),
		grpc.WithChainStreamInterceptor(
			otelgrpc.StreamClientInterceptor(),
			grpc_prometheus.StreamClientInterceptor,
		),
	)
	if err != nil {
		c.in.Logger.Critical(ctx, errors.New("could not connect to grpc server: %s", err.Error()))
		return nil, err
	}

	return conn, nil
}
