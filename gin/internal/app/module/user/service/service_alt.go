package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/dto"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/entity"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/repository"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/schema"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/adapter"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/data/nosql/mongodb"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/time"

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

func (s *ServiceAlt) Create(d *dto.CreateUpdateUserDto, ctx *gin.Context) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	res, err := s.repository.Create(ctx, &v, nil)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), ctx)
}

func (s *ServiceAlt) ReadMany(limit, skip int, ctx *gin.Context) {
	opts := mongodb.BuildPaginatedOpts(limit, skip)
	c, err := s.repository.ReadMany(
		ctx,
		bson.D{},
		&opts,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	var data []schema.User
	data, err = mongodb.DecodeCursor[[]schema.User](c, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), ctx)
			return
		} else if err == mongo.ErrNilCursor {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), ctx)
			return
		}
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	if data == nil {
		data = []schema.User{}
	}
	m := make(map[string]any)
	m["items"] = data
	m["limit"] = limit
	m["page"] = skip
	response.Respond(http.StatusOK, response.BuildData(m), ctx)
}

func (s *ServiceAlt) ReadManyWithNestedDocQuery(limit, skip int, key0, key1 string, ctx *gin.Context) {
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
		ctx,
		filter,
		&opts,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	var data []schema.User
	data, err = mongodb.DecodeCursor[[]schema.User](c, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), ctx)
			return
		} else if err == mongo.ErrNilCursor {
			// This error means your query did not match any documents.
			response.Respond(http.StatusOK, make([]any, 0), ctx)
			return
		}
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	if data == nil {
		data = []schema.User{}
	}
	response.Respond(http.StatusOK, response.BuildData(data), ctx)
}

func (s *ServiceAlt) ReadOne(id string, ctx *gin.Context) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	res := s.repository.ReadOne(
		ctx,
		filter,
		nil,
	)
	var data schema.User
	data, err = mongodb.DecodeSingleResult[schema.User](res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.

		}
		response.RespondError(http.StatusNotFound, err, ctx)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(data), ctx)
}

func (s *ServiceAlt) Update(id string, d *dto.CreateUpdateUserDto, ctx *gin.Context) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	doc := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: d.Name}, {Key: "updatedAt", Value: time.Now()}}}}
	res, err := s.repository.Update(
		ctx,
		filter,
		doc,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), ctx)
}

func (s *ServiceAlt) Delete(id string, ctx *gin.Context) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	res, err := s.repository.Delete(ctx, filter, nil)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), ctx)
}
