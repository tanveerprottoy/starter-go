package user

import (
	"context"
	"log"
	"net/http"
	"time"
	"txp/restapistarter/app/module/user/dto"
	"txp/restapistarter/app/module/user/repository"
	"txp/restapistarter/app/module/user/schema"
	"txp/restapistarter/app/util"
	"txp/restapistarter/pkg/coreutil"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/responseutil"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoService struct {
	repository *repository.UserMongoRepository
}

func NewUserMongoService(repository *repository.UserMongoRepository) *UserMongoService {
	s := new(UserMongoService)
	s.repository = repository
	return s
}

func (s *UserMongoService) Create(w http.ResponseWriter, r *http.Request) {
	var b *dto.CreateUpdateUserDto
	err := coreutil.Decode(r, b)
	if err != nil {
		responseutil.RespondError(http.StatusBadRequest, err, w)
		return
	}
	res, err := s.repository.Create(
		util.UsersCollection,
		r.Context(),
		&schema.UserSchema{
			Name: b.Name,
		},
		nil,
	)
	if err != nil {
		responseutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	log.Println(res)
	responseutil.Respond(http.StatusOK, res, w)
}

func (s *UserMongoService) ReadMany(w http.ResponseWriter, r *http.Request) {
	// context.TODO()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"name", bson.D{{"$eq", "a"}}}}
	c, err := s.repository.ReadMany(
		util.UsersCollection,
		ctx,
		filter,
		nil,
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
		} else if err == mongo.ErrNilCursor {
			// This error means your query did not match any documents.
		}
		responseutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	data, err := mongodb.Decode(c)
	if err != nil {
		responseutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	responseutil.Respond(http.StatusOK, data, w)
}

func (s *UserMongoService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := coreutil.GetURLParam(r, util.UrlKeyId)
	filter := bson.D{{"_id", bson.D{{"$eq", userId}}}}
	res := s.repository.ReadOne(
		util.UsersCollection,
		r.Context(),
		filter,
		nil,
	)
	responseutil.Respond(http.StatusOK, res, w)
}

func (s *UserMongoService) Update(w http.ResponseWriter, r *http.Request) {
	userId := coreutil.GetURLParam(r, util.UrlKeyId)
	filter := bson.D{{"_id", bson.D{{"$eq", userId}}}}
	var b *dto.CreateUpdateUserDto
	err := coreutil.Decode(r, b)
	if err != nil {
		responseutil.RespondError(http.StatusBadRequest, err, w)
		return
	}
	res, err := s.repository.Update(
		util.UsersCollection,
		r.Context(),
		filter,
		&schema.UserSchema{
			Name: b.Name,
		},
		nil,
	)
	if err != nil {
		responseutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	log.Println(res)
	responseutil.Respond(http.StatusOK, res, w)
}

func (s *UserMongoService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := coreutil.GetURLParam(r, util.UrlKeyId)
	filter := bson.D{{"_id", bson.D{{"$eq", userId}}}}
	res, err := s.repository.Delete(
		util.UsersCollection,
		r.Context(),
		filter,
		nil,
	)
	if err != nil {
		responseutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	log.Println(res)
	responseutil.Respond(http.StatusOK, res, w)
}
