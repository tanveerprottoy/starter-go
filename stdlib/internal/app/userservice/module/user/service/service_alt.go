package service

import (
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/dto"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/repository"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/schema"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/adapter"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/data/nosql/mongodb"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/timepkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceAlt struct {
	repository *repository.RepositoryAlt
}

func NewServiceAlt(r *repository.RepositoryAlt) *ServiceAlt {
	s := new(ServiceAlt)
	s.repository = r
	return s
}

func (s *ServiceAlt) Create(d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	n := timepkg.NowUnixMilli()
	v.CreatedAt = n
	v.UpdatedAt = n
	res, err := s.repository.Create(
		r.Context(),
		&v,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), w)
}

func (s *ServiceAlt) ReadMany(limit, skip int, w http.ResponseWriter, r *http.Request) {
	opts := mongodb.BuildPaginatedOpts(limit, skip)
	c, err := s.repository.ReadMany(
		r.Context(),
		bson.D{},
		&opts,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	var data []schema.User
	data, err = mongodb.DecodeCursor[[]schema.User](c, r.Context())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), w)
			return
		} else if err == mongo.ErrNilCursor {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), w)
			return
		}
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	if data == nil {
		data = []schema.User{}
	}
	m := make(map[string]any)
	m["items"] = data
	m["limit"] = limit
	m["page"] = skip
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (s *ServiceAlt) ReadManyWithNestedDocQuery(limit, skip int, key0, key1 string, w http.ResponseWriter, r *http.Request) {
	opts := mongodb.BuildPaginatedOpts(limit, skip)
	filter := bson.D{}
	if key0 != "" {
		filter = bson.D{
			{Key: "addresses.text", Value: key0},
		}
	} else if key1 != "" {
		filter = bson.D{
			{Key: "phones.text", Value: key1},
		}
	}
	c, err := s.repository.ReadMany(
		r.Context(),
		filter,
		&opts,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	var data []schema.User
	data, err = mongodb.DecodeCursor[[]schema.User](c, r.Context())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), w)
			return
		} else if err == mongo.ErrNilCursor {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), w)
			return
		}
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	if data == nil {
		data = []schema.User{}
	}
	response.Respond(http.StatusOK, response.BuildData(data), w)
}

func (s *ServiceAlt) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	res := s.repository.ReadOne(
		r.Context(),
		filter,
		nil,
	)
	var data schema.User
	data, err = mongodb.DecodeSingleResult[schema.User](res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.

		}
		response.RespondError(http.StatusNotFound, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(data), w)
}

func (s *ServiceAlt) Update(id string, d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	doc := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: d.Name}, {Key: "updatedAt", Value: timepkg.NowUnixMilli()}}}}
	res, err := s.repository.Update(
		r.Context(),
		filter,
		doc,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), w)
}

func (s *ServiceAlt) Delete(id string, w http.ResponseWriter, r *http.Request) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	res, err := s.repository.Delete(
		r.Context(),
		filter,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), w)
}
