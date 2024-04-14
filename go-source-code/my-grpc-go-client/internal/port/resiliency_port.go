package port

import (
	"context"

	resl "github.com/timpamungkas/my-grpc-proto/protogen/go/resiliency"
	"google.golang.org/grpc"
)

type ResiliencyClientPort interface {
	UnaryResiliency(ctx context.Context, in *resl.ResiliencyRequest,
		opts ...grpc.CallOption) (*resl.ResiliencyResponse, error)
	ServerStreamingResiliency(ctx context.Context, in *resl.ResiliencyRequest,
		opts ...grpc.CallOption) (resl.ResiliencyService_ServerStreamingResiliencyClient, error)
	ClientStreamingResiliency(ctx context.Context,
		opts ...grpc.CallOption) (resl.ResiliencyService_ClientStreamingResiliencyClient, error)
	BiDirectionalResiliency(ctx context.Context,
		opts ...grpc.CallOption) (resl.ResiliencyService_BiDirectionalResiliencyClient, error)
}

type ResiliencyWithMetadataClientPort interface {
	UnaryResiliencyWithMetadata(ctx context.Context,
		in *resl.ResiliencyRequest, opts ...grpc.CallOption) (*resl.ResiliencyResponse, error)
	ServerStreamingResiliencyWithMetadata(ctx context.Context,
		in *resl.ResiliencyRequest, opts ...grpc.CallOption) (
		resl.ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadataClient, error)
	ClientStreamingResiliencyWithMetadata(ctx context.Context,
		opts ...grpc.CallOption) (
		resl.ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadataClient, error)
	BiDirectionalResiliencyWithMetadata(ctx context.Context, opts ...grpc.CallOption) (
		resl.ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadataClient, error)
}
