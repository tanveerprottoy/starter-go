package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/repository"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/grpcpkg"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ServiceRPC struct {
	repository *repository.RepositoryRPC
}

func NewServiceRPC(r *repository.RepositoryRPC) *ServiceRPC {
	s := new(ServiceRPC)
	s.repository = r
	return s
}

func (s *ServiceRPC) Create(ctx context.Context, u *proto.User) (*proto.User, error) {
	lastId, err := s.repository.Create(
		u,
	)
	if err != nil || lastId != "" {
		return nil, grpcpkg.RespondError(
			codes.Unknown,
			constant.UnknownError,
		)
	}
	return u, nil
}

func (s *ServiceRPC) ReadMany(ctx context.Context, v *proto.VoidParam) (*proto.Users, error) {
	log.Print("ReadMany rpc")
	d := &proto.Users{}
	rows, err := s.repository.ReadMany()
	if err != nil {
		return nil, grpcpkg.RespondError(
			codes.Unknown,
			constant.UnknownError,
		)
	}
	var (
		users      []*proto.User
		id         string
		name       string
		created_at time.Time
		updated_at time.Time
	)
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&id, &name, &created_at, &updated_at); err != nil {
			return nil, fmt.Errorf("ReadMany %v", err)
		}
		users = append(users, &proto.User{
			Id:   id,
			Name: name,
			CreatedAt: timestamppb.New(
				created_at,
			),
			UpdatedAt: timestamppb.New(
				updated_at,
			),
		})
	}
	d.Users = users
	return d, err
}

/* func (s *ServiceRPC) ReadUserStream(
	v *proto.VoidParam,
	serv proto.ServiceRPC_ReadUserStreamServer,
) (*proto.Users, error) {
	return &proto.Users{}, nil
} */

func (s *ServiceRPC) ReadOne(ctx context.Context, strVal *wrapperspb.StringValue) (*proto.User, error) {
	row := s.repository.ReadOne(
		strVal.Value,
	)
	if row == nil {
		return nil, grpcpkg.RespondError(
			codes.NotFound,
			constant.NotFound,
		)
	}
	var (
		id         string
		name       string
		created_at time.Time
		updated_at time.Time
	)
	if err := row.Scan(&id, &name, &created_at, &updated_at); err != nil {
		return nil, fmt.Errorf("ReadOne %v", err)
	}
	u := &proto.User{
		Id:   id,
		Name: name,
		CreatedAt: timestamppb.New(
			created_at,
		),
		UpdatedAt: timestamppb.New(
			updated_at,
		),
	}
	return u, nil
}

func (s *ServiceRPC) Update(ctx context.Context, p *proto.UpdateUserParam) (*proto.User, error) {
	r, err := s.repository.Update(
		p.Id,
		p.User,
	)
	if err != nil || r <= 0 {
		return nil, grpcpkg.RespondError(
			codes.Unknown,
			constant.UnknownError,
		)
	}
	return p.User, nil
}

func (s *ServiceRPC) Delete(ctx context.Context, strVal *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	r, err := s.repository.Delete(
		strVal.Value,
	)
	if err != nil || r <= 0 {
		return nil, grpcpkg.RespondError(
			codes.Unknown,
			constant.UnknownError,
		)
	}
	return &wrapperspb.BoolValue{Value: true}, nil
}
