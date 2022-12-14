package service

import (
	"net/http"
	"txp/restapistarter/internal/app/module/user/dto"
	"txp/restapistarter/internal/app/module/user/repository"
	"txp/restapistarter/internal/app/module/user/schema"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoService struct {
	repository *repository.UserMongoRepository
}

func NewUserMongoService(r *repository.UserMongoRepository) *UserMongoService {
	s := new(UserMongoService)
	s.repository = r
	return s
}

func (s *UserMongoService) Create(p []byte, w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BytesToValue[schema.User](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	res, err := s.repository.Create(
		r.Context(),
		&d,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), w)
}

func (s *UserMongoService) ReadMany(limit, skip int, w http.ResponseWriter, r *http.Request) {
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

func (s *UserMongoService) ReadManyWithNestedDocQuery(limit, skip int, key0, key1 string, w http.ResponseWriter, r *http.Request) {
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

func (s *UserMongoService) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
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

func (s *UserMongoService) Update(id string, p []byte, w http.ResponseWriter, r *http.Request) {
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d, err := adapter.BytesToValue[dto.CreateUpdateUserDto](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	doc := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: d.Name}, {Key: "updatedAt", Value: time.Now()}}}}
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

func (s *UserMongoService) Delete(id string, w http.ResponseWriter, r *http.Request) {
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
