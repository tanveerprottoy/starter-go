package user

import (
	"log"
	"net/http"
	"txp/restapistarter/internal/app/module/user/repository"
	"txp/restapistarter/internal/app/module/user/schema"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/router"
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
		constant.UsersCollection,
		r.Context(),
		d,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(res), w)
}

func (s *UserMongoService) ReadMany(w http.ResponseWriter, r *http.Request) {
	// context.TODO()
	/* ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() */
	// filter := bson.D{{"name", bson.D{{"$eq", "a"}}}}
	/* l := int64(limit)
	skip := int64(page*limit - limit) */
	// opts := options.FindOptions{Limit: &l, Skip: &skip}
	c, err := s.repository.ReadMany(
		constant.UsersCollection,
		r.Context(),
		bson.D{},
		nil,
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
	response.Respond(http.StatusOK, response.BuildData(data), w)
}

func (s *UserMongoService) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	objId, err := mongodb.BuildObjectID(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	res := s.repository.ReadOne(
		constant.UsersCollection,
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
	d, err := adapter.BytesToValue[schema.User](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: objId}}}}
	doc := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: d.Name}}}}
	res, err := s.repository.Update(
		constant.UsersCollection,
		r.Context(),
		filter,
		doc,
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	log.Println(res)
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
		constant.UsersCollection,
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
