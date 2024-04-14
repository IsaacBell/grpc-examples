package hello

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/timpamungkas/my-grpc-go-client/internal/port"
	"github.com/timpamungkas/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

type HelloAdapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAdapter, error) {
	client := hello.NewHelloServiceClient(conn)

	return &HelloAdapter{
		helloClient: client,
	}, nil
}

func (a *HelloAdapter) SayHello(ctx context.Context, name string) (*hello.HelloResponse, error) {
	helloRequest := &hello.HelloRequest{
		Name: name,
	}

	greet, err := a.helloClient.SayHello(ctx, helloRequest)

	if err != nil {
		log.Fatalln("Error on SayHello :", err)
	}

	return greet, nil
}

func (a *HelloAdapter) SayManyHellos(ctx context.Context, name string) {
	helloRequest := &hello.HelloRequest{
		Name: name,
	}

	greetStream, err := a.helloClient.SayManyHellos(ctx, helloRequest)

	if err != nil {
		log.Fatalln("Error on SayManyHellos :", err)
	}

	for {
		greet, err := greetStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error on SayManyHellos :", err)
		}

		log.Println(greet.Greet)
	}
}

func (a *HelloAdapter) SayHelloToEveryone(ctx context.Context, names []string) {
	greetStream, err := a.helloClient.SayHelloToEveryone(ctx)

	if err != nil {
		log.Fatalln("Error on SayHelloToEveryone :", err)
	}

	for _, name := range names {
		req := &hello.HelloRequest{
			Name: name,
		}

		greetStream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	res, err := greetStream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Error on SayHelloToEveryone :", err)
	}

	log.Println(res.Greet)
}

func (a *HelloAdapter) SayHelloContinuous(ctx context.Context, names []string) {
	greetStream, err := a.helloClient.SayHelloContinuous(ctx)

	if err != nil {
		log.Fatalln("Error on SayHelloContinuous :", err)
	}

	greetChan := make(chan struct{})

	go func() {
		for _, name := range names {
			req := &hello.HelloRequest{
				Name: name,
			}

			greetStream.Send(req)
		}

		greetStream.CloseSend()
	}()

	go func() {
		for {
			greet, err := greetStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalln("Error on SayHelloContinuous :", err)
			}

			log.Println(greet.Greet)
		}

		close(greetChan)
	}()

	<-greetChan
}
