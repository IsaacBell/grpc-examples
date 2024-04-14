package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/timpamungkas/my-grpc-go-server/internal/port"
	"github.com/timpamungkas/my-grpc-proto/protogen/go/bank"
	"github.com/timpamungkas/my-grpc-proto/protogen/go/hello"
	resl "github.com/timpamungkas/my-grpc-proto/protogen/go/resiliency"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService      port.HelloServicePort
	bankService       port.BankServicePort
	resiliencyService port.ResiliencyServicePort
	grpcPort          int
	server            *grpc.Server
	hello.HelloServiceServer
	bank.BankServiceServer
	resl.ResiliencyServiceServer
	resl.ResiliencyWithMetadataServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, bankService port.BankServicePort,
	resiliencyService port.ResiliencyServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService:      helloService,
		bankService:       bankService,
		resiliencyService: resiliencyService,
		grpcPort:          grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n", a.grpcPort)

	// creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")

	// if err != nil {
	// 	log.Fatalln("Can't create server credentials :", err)
	// }

	grpcServer := grpc.NewServer(
	// grpc.Creds(creds),
	// grpc.ChainUnaryInterceptor(
	// 	interceptor.LogUnaryServerInterceptor(),
	// 	interceptor.BasicUnaryServerInterceptor(),
	// ),
	// grpc.ChainStreamInterceptor(
	// 	interceptor.LogStreamServerInterceptor(),
	// 	interceptor.BasicStreamServerInterceptor(),
	// ),
	)

	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)
	bank.RegisterBankServiceServer(grpcServer, a)
	resl.RegisterResiliencyServiceServer(grpcServer, a)
	resl.RegisterResiliencyWithMetadataServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
