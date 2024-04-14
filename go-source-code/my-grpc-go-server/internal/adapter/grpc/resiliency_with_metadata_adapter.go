package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	resl "github.com/timpamungkas/my-grpc-proto/protogen/go/resiliency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func dummyRequestMetadata(ctx context.Context) {
	if requestMetadata, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println("Request metadata :")
		for k, v := range requestMetadata {
			log.Printf("  %v : %v\n", k, v)
		}
	} else {
		log.Println("Request metadata not found")
	}
}

func dummyResponseMetadata() metadata.MD {
	md := map[string]string{
		"grpc-server-time":     fmt.Sprint(time.Now().Format("15:04:05")),
		"grpc-server-location": "Jakarta, Indonesia",
		"grpc-response-uuid":   uuid.New().String(),
	}

	return metadata.New(md)
}

func (a *GrpcAdapter) UnaryResiliencyWithMetadata(ctx context.Context, req *resl.ResiliencyRequest) (
	*resl.ResiliencyResponse, error) {
	log.Println("UnaryResiliencyWithMetadata called")
	str, sts := a.resiliencyService.GenerateResiliency(req.MinDelaySecond,
		req.MaxDelaySecond, req.StatusCodes)

	// read request metadata
	dummyRequestMetadata(ctx)

	if errStatus := generateErrStatus(sts); errStatus != nil {
		return nil, errStatus
	}

	// add response metadata
	grpc.SendHeader(ctx, dummyResponseMetadata())

	return &resl.ResiliencyResponse{
		DummyString: str,
	}, nil
}

func (a *GrpcAdapter) ServerStreamingResiliencyWithMetadata(req *resl.ResiliencyRequest,
	stream resl.ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadataServer) error {
	log.Println("ServerStreamingResiliencyWithMetadata called")

	context := stream.Context()

	dummyRequestMetadata(context)

	if err := stream.SendHeader(dummyResponseMetadata()); err != nil {
		log.Println("Error while sending response metadata :", err)
	}

	for {
		select {
		case <-context.Done():
			log.Println("Client cancelled request")
			return nil
		default:
			str, sts := a.resiliencyService.GenerateResiliency(req.MinDelaySecond,
				req.MaxDelaySecond, req.StatusCodes)

			if errStatus := generateErrStatus(sts); errStatus != nil {
				return errStatus
			}

			stream.Send(&resl.ResiliencyResponse{
				DummyString: str,
			})
		}
	}
}

func (a *GrpcAdapter) ClientStreamingResiliencyWithMetadata(
	stream resl.ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadataServer) error {
	log.Println("ClientStreamingResiliencyWithMetadata called")

	i := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res := resl.ResiliencyResponse{
				DummyString: fmt.Sprintf("Received %v requests from client", strconv.Itoa(i)),
			}

			if err := stream.SendHeader(dummyResponseMetadata()); err != nil {
				log.Println("Error while sending response metadata :", err)
			}

			return stream.SendAndClose(&res)
		}

		context := stream.Context()
		dummyRequestMetadata(context)

		if req != nil {
			_, sts := a.resiliencyService.GenerateResiliency(req.MinDelaySecond,
				req.MaxDelaySecond, req.StatusCodes)

			if errStatus := generateErrStatus(sts); errStatus != nil {
				return errStatus
			}
		}

		i = i + 1
	}
}

func (a *GrpcAdapter) BiDirectionalResiliencyWithMetadata(
	stream resl.ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadataServer) error {
	log.Println("BiDirectionalResiliencyWithMetadata called")

	context := stream.Context()

	if err := stream.SendHeader(dummyResponseMetadata()); err != nil {
		log.Println("Error while sending response metadata :", err)
	}

	for {
		select {
		case <-context.Done():
			log.Println("Client cancelled request")
			return nil
		default:
			req, err := stream.Recv()

			if err == io.EOF {
				return nil
			}

			if err != nil {
				log.Fatalln("Error while reading from client :", err)
			}

			dummyRequestMetadata(context)

			str, sts := a.resiliencyService.GenerateResiliency(req.MinDelaySecond,
				req.MaxDelaySecond, req.StatusCodes)

			if errStatus := generateErrStatus(sts); errStatus != nil {
				return errStatus
			}

			err = stream.Send(
				&resl.ResiliencyResponse{
					DummyString: str,
				},
			)

			if err != nil {
				log.Fatalln("Error while sending response to client :", err)
			}
		}
	}
}
