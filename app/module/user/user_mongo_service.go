package user

import (
	"context"
	"net/http"
	"time"
	"txp/restapistarter/app/module/user/entity"
	"txp/restapistarter/app/module/user/repository"
	"txp/restapistarter/app/util"
	"txp/restapistarter/pkg/coreutil"
	"txp/restapistarter/pkg/data/nosql/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoService struct {
	repository *repository.UserMongoRepository[entity.UserSchema]
}

func NewUserMongoService(repository *repository.UserMongoRepository[entity.UserSchema]) *UserMongoService {
	s := new(UserMongoService)
	s.repository = repository
	return s
}

func (s *UserMongoService) Create(w http.ResponseWriter, r *http.Request) {
	/* var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		coreutil.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		&entity.User{
			Name: b.Name,
		},
	)
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	coreutil.Respond(http.StatusCreated, b, w) */
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
		coreutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	data, err := mongodb.Decode(c)
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	coreutil.Respond(http.StatusOK, data, w)
}

func (s *UserMongoService) ReadOne(w http.ResponseWriter, r *http.Request) {
	/* userId := chi.URLParam(r, util.UrlKeyId)
	row := s.repository.ReadOne(userId)
	if row == nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	e := new(entity.User)
	d, err := sqlUtil.GetEntity(
		row,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	coreutil.Respond(http.StatusOK, d, w) */
}

func (s *UserMongoService) Update(w http.ResponseWriter, r *http.Request) {
	/* userId := chi.URLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		coreutil.RespondError(
			http.StatusBadRequest,
			err,
			w,
		)
		return
	}
	rowsAffected, err := s.repository.Update(
		userId,
		&entity.User{
			Name: b.Name,
		},
	)
	if err != nil || rowsAffected <= 0 {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	coreutil.Respond(http.StatusOK, b, w) */
}

func (s *UserMongoService) Delete(w http.ResponseWriter, r *http.Request) {
	/* userId := chi.URLParam(r, util.UrlKeyId)
	rowsAffected, err := s.repository.Delete(userId)
	if err != nil || rowsAffected <= 0 {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	coreutil.Respond(
		http.StatusOK,
		map[string]bool{"success": true},
		w,
	) */
}
