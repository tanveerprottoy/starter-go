package rpc

import (
	"context"
	"txp/userservice/app/module/user/proto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RPC struct {
	proto.UnimplementedUserServiceServer
	service *ServiceRPC
}

func NewRPC(service *ServiceRPC) *RPC {
	s := new(RPC)
	s.service = service
	return s
}

func (h *UserRPC) CreateUser(
	ctx context.Context,
	u *proto.User,
) (*proto.User, error) {
	return h.service.Create(
		ctx,
		u,
	)
}

func (h *UserRPC) ReadUsers(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Users, error) {
	return h.service.ReadMany(
		ctx,
		v,
	)
}

/* func (h *UserRPC) ReadUserStream(
	v *proto.VoidParam,
	serv proto.UserService_ReadUserStreamServer,
) (*proto.Users, error) {
	return nil, nil
	h.service.ReadMany(
		ctx,
		v,
	)
} */

func (h *UserRPC) ReadUser(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*proto.User, error) {
	return h.service.ReadOne(
		ctx,
		strVal,
	)
}

func (h *UserRPC) UpdateUser(
	ctx context.Context,
	p *proto.UpdateUserParam,
) (*proto.User, error) {
	return h.service.Update(
		ctx,
		p,
	)
}

func (h *UserRPC) DeleteUser(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*wrapperspb.BoolValue, error) {
	return h.service.Delete(
		ctx,
		strVal,
	)
}
