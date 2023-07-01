package grpcpkg

import (
	"log"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/proto"
	"google.golang.org/grpc"
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
	Conns[0], err = grpc.Dial(
		"0.0.0.0:5000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("userservice did not connect: %v", err)
	}
	Conns[1], err = grpc.Dial(
		"0.0.0.0:5001",
		grpc.WithInsecure(),
		grpc.WithBlock(),
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
