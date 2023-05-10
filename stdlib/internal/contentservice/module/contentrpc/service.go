package content

import (
	"context"
	"fmt"
	"log"
	"time"
	"txp/contentservice/app/module/content/proto"
	"txp/contentservice/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {
	s := new(Service)
	s.repository = repository
	return s
}

func (s *Service) Create(
	ctx context.Context,
	u *proto.Content,
) (*proto.Content, error) {
	l, err := s.repository.Create(
		u,
	)
	if err != nil || l != "" {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return u, nil
}

func (s *Service) ReadMany(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Contents, error) {
	log.Print("ReadMany rpc")
	d := &proto.Contents{}
	rows, err := s.repository.ReadMany()
	var (
		contents   []*proto.Content
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
		contents = append(contents, &proto.Content{
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
	d.Contents = contents
	return d, err
}

/* func (s *ContentService) Readcontentstream(
	v *proto.VoidParam,
	serv proto.ContentService_ReadcontentstreamServer,
) (*proto.Contents, error) {
	return &proto.Contents{}, nil
} */

func (s *Service) ReadOne(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*proto.Content, error) {
	row := s.repository.ReadOne(
		strVal.Value,
	)
	if row == nil {
		return nil, util.RespondError(
			codes.NotFound,
			util.NotFound,
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
	u := &proto.Content{
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

func (s *Service) Update(
	ctx context.Context,
	p *proto.UpdateContentParam,
) (*proto.Content, error) {
	r, err := s.repository.Update(
		p.Id,
		p.Content,
	)
	if err != nil || r <= 0 {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return p.Content, nil
}

func (s *Service) Delete(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*wrapperspb.BoolValue, error) {
	r, err := s.repository.Delete(
		strVal.Value,
	)
	if err != nil || r <= 0 {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return &wrapperspb.BoolValue{Value: true}, nil
}
