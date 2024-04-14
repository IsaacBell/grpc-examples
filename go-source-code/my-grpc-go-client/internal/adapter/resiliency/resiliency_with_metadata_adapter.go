package resiliency

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/google/uuid"
	resl "github.com/timpamungkas/my-grpc-proto/protogen/go/resiliency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func sampleRequestMetadata() metadata.MD {
	md := map[string]string{
		"grpc-client-time":  fmt.Sprintf(time.Now().Format("15:04:05")),
		"grpc.client-os":    runtime.GOOS,
		"grpc-request-uuid": uuid.New().String(),
	}
	return metadata.New(md)
}

func sampleResponseMetadata(md metadata.MD) {
	if md.Len() == 0 {
		log.Println("Response metadata not found")
	} else {
		log.Println("Response metadata :")
		for k, v := range md {
			log.Printf("  %v : %v\n", k, v)
		}
	}
}

func (a *ResiliencyAdapter) UnaryResiliencyWithMetadata(ctx context.Context, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32) (*resl.ResiliencyResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, sampleRequestMetadata())

	resiliencyRequest := &resl.ResiliencyRequest{
		MinDelaySecond: minDelaySecond,
		MaxDelaySecond: maxDelaySecond,
		StatusCodes:    statusCodes,
	}

	var responseMetadata metadata.MD
	res, err := a.resiliencyWithMetadataClient.UnaryResiliencyWithMetadata(ctx,
		resiliencyRequest, grpc.Header(&responseMetadata))

	if err != nil {
		log.Println("Error on UnaryResiliencyWithMetadata :", err)
		return nil, err
	}

	sampleResponseMetadata(responseMetadata)

	return res, nil
}

func (a *ResiliencyAdapter) ServerStreamingResiliencyWithMetadata(ctx context.Context, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32) {
	ctx = metadata.NewOutgoingContext(ctx, sampleRequestMetadata())
	resiliencyRequest := &resl.ResiliencyRequest{
		MinDelaySecond: minDelaySecond,
		MaxDelaySecond: maxDelaySecond,
		StatusCodes:    statusCodes,
	}

	reslStream, err := a.resiliencyWithMetadataClient.ServerStreamingResiliencyWithMetadata(
		ctx, resiliencyRequest)

	if err != nil {
		log.Fatalln("Error on ServerStreamingResiliencyWithMetadata :", err)
	}

	if responseMetadata, err := reslStream.Header(); err == nil {
		sampleResponseMetadata(responseMetadata)
	}

	for {
		res, err := reslStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error on ServerStreamingResiliencyWithMetadata :", err)
		}

		log.Println(res.DummyString)
	}
}

func (a *ResiliencyAdapter) ClientStreamingResiliencyWithMetadata(ctx context.Context, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, count int) {
	reslStream, err := a.resiliencyWithMetadataClient.ClientStreamingResiliencyWithMetadata(ctx)

	if err != nil {
		log.Fatalln("Error on ClientStreamingResiliencyWithMetadata :", err)
	}

	for i := 1; i <= count; i++ {
		ctx = metadata.NewOutgoingContext(ctx, sampleRequestMetadata())

		resiliencyRequest := &resl.ResiliencyRequest{
			MinDelaySecond: minDelaySecond,
			MaxDelaySecond: maxDelaySecond,
			StatusCodes:    statusCodes,
		}

		reslStream.Send(resiliencyRequest)
	}

	res, err := reslStream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Error on ClientStreamingResiliencyWithMetadata :", err)
	}

	if responseMetadata, err := reslStream.Header(); err == nil {
		sampleResponseMetadata(responseMetadata)
	}

	log.Println(res.DummyString)
}

func (a *ResiliencyAdapter) BiDirectionalResiliencyWithMetadata(ctx context.Context, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, count int) {
	reslStream, err := a.resiliencyWithMetadataClient.BiDirectionalResiliencyWithMetadata(ctx)

	if err != nil {
		log.Fatalln("Error on BiDirectionalResiliencyWithMetadata :", err)
	}

	if responseMetadata, err := reslStream.Header(); err == nil {
		sampleResponseMetadata(responseMetadata)
	}

	reslChan := make(chan struct{})

	go func() {
		for i := 1; i <= count; i++ {
			ctx = metadata.NewOutgoingContext(ctx, sampleRequestMetadata())

			resiliencyRequest := &resl.ResiliencyRequest{
				MinDelaySecond: minDelaySecond,
				MaxDelaySecond: maxDelaySecond,
				StatusCodes:    statusCodes,
			}

			reslStream.Send(resiliencyRequest)
		}

		reslStream.CloseSend()
	}()

	go func() {
		for {
			res, err := reslStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalln("Error on BiDirectionalResiliencyWithMetadata :", err)
			}

			log.Println(res.DummyString)
		}

		close(reslChan)
	}()

	<-reslChan
}
