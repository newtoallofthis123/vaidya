// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: transcribe.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AudioService_TranscribeAudio_FullMethodName = "/AudioService/TranscribeAudio"
)

// AudioServiceClient is the client API for AudioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioServiceClient interface {
	TranscribeAudio(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AudioFile, TranscribeResponse], error)
}

type audioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioServiceClient(cc grpc.ClientConnInterface) AudioServiceClient {
	return &audioServiceClient{cc}
}

func (c *audioServiceClient) TranscribeAudio(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AudioFile, TranscribeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AudioService_ServiceDesc.Streams[0], AudioService_TranscribeAudio_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AudioFile, TranscribeResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AudioService_TranscribeAudioClient = grpc.ClientStreamingClient[AudioFile, TranscribeResponse]

// AudioServiceServer is the server API for AudioService service.
// All implementations must embed UnimplementedAudioServiceServer
// for forward compatibility.
type AudioServiceServer interface {
	TranscribeAudio(grpc.ClientStreamingServer[AudioFile, TranscribeResponse]) error
	mustEmbedUnimplementedAudioServiceServer()
}

// UnimplementedAudioServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAudioServiceServer struct{}

func (UnimplementedAudioServiceServer) TranscribeAudio(grpc.ClientStreamingServer[AudioFile, TranscribeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method TranscribeAudio not implemented")
}
func (UnimplementedAudioServiceServer) mustEmbedUnimplementedAudioServiceServer() {}
func (UnimplementedAudioServiceServer) testEmbeddedByValue()                      {}

// UnsafeAudioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioServiceServer will
// result in compilation errors.
type UnsafeAudioServiceServer interface {
	mustEmbedUnimplementedAudioServiceServer()
}

func RegisterAudioServiceServer(s grpc.ServiceRegistrar, srv AudioServiceServer) {
	// If the following call pancis, it indicates UnimplementedAudioServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AudioService_ServiceDesc, srv)
}

func _AudioService_TranscribeAudio_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioServiceServer).TranscribeAudio(&grpc.GenericServerStream[AudioFile, TranscribeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AudioService_TranscribeAudioServer = grpc.ClientStreamingServer[AudioFile, TranscribeResponse]

// AudioService_ServiceDesc is the grpc.ServiceDesc for AudioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AudioService",
	HandlerType: (*AudioServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TranscribeAudio",
			Handler:       _AudioService_TranscribeAudio_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transcribe.proto",
}
