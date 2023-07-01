package rpc

import (
	"context"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/service"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RPC struct {
	proto.UnimplementedUserServiceServer
	service *service.ServiceRPC
}

func NewRPC(service *service.ServiceRPC) *RPC {
	s := new(RPC)
	s.service = service
	return s
}

func (h *RPC) CreateUser(ctx context.Context, u *proto.User) (*proto.User, error) {
	return h.service.Create(ctx, u)
}

func (h *RPC) ReadUsers(ctx context.Context, v *proto.VoidParam) (*proto.Users, error) {
	return h.service.ReadMany(ctx, v)
}

/* func (h *service.ServiceRPC) ReadUserStream(v *proto.VoidParam, srv proto.UserService_ReadUserStreamServer) (*proto.Users, error) {
	return nil, nil
	h.service.ReadMany(ctx,v)
} */

func (h *RPC) ReadUser(ctx context.Context, strVal *wrapperspb.StringValue) (*proto.User, error) {
	return h.service.ReadOne(ctx, strVal)
}

func (h *RPC) UpdateUser(ctx context.Context, p *proto.UpdateUserParam) (*proto.User, error) {
	return h.service.Update(ctx, p)
}

func (h *RPC) DeleteUser(ctx context.Context, strVal *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return h.service.Delete(ctx, strVal)
}
