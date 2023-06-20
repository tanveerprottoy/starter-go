package grpcpkg

import (
	"log"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/proto"
	grpc "google.golang.org/grpc"
)

var (
	Conns                [2]*grpc.ClientConn
	UserServiceClient    proto.UserServiceClient
	ContentServiceClient _contentProto.ContentServiceClient
)

func init() {
	initClientConns()
	initServiceClients()
}

func initClientConns() {
	var err error
	Conns[0], err = _grpc.Dial(
		"0.0.0.0:5000",
		_grpc.WithInsecure(),
		_grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("userservice did not connect: %v", err)
	}
	Conns[1], err = _grpc.Dial(
		"0.0.0.0:5001",
		_grpc.WithInsecure(),
		_grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("contentservice did not connect: %v", err)
	}
}

func initServiceClients() {
	UserServiceClient = proto.NewUserServiceClient(
		Conns[0],
	)
	ContentServiceClient = _contentProto.NewContentServiceClient(
		Conns[1],
	)
}
