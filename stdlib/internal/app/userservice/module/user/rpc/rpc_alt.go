package rpc

import (
	"context"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/service"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RPCAlt struct {
	proto.UnimplementedUserServiceServer
	service *service.ServiceRPC
}

func NewRPCAlt(service *service.ServiceRPC) *RPCAlt {
	s := new(RPCAlt)
	s.service = service
	return s
}

func (h *RPCAlt) CreateUser(ctx context.Context, u *proto.User) (*proto.User, error) {
	return h.service.Create(ctx, u)
}

func (h *RPCAlt) ReadUsers(ctx context.Context, v *proto.VoidParam) (*proto.Users, error) {
	return h.service.ReadMany(ctx, v)
}

/* func (h *RPCAlt) ReadUserStream(v *proto.VoidParam,serv proto.UserService_ReadUserStreamServer) (*proto.Users, error) {
	return nil, nil
	h.service.ReadMany(ctx,v)
} */

func (h *RPCAlt) ReadUser(ctx context.Context, strVal *wrapperspb.StringValue) (*proto.User, error) {
	return h.service.ReadOne(ctx, strVal)
}

func (h *RPCAlt) UpdateUser(ctx context.Context, p *proto.UpdateUserParam) (*proto.User, error) {
	return h.service.Update(ctx, p)
}

func (h *RPCAlt) DeleteUser(ctx context.Context, strVal *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return h.service.Delete(ctx, strVal)
}
