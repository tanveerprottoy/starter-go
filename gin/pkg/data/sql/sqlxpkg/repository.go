package sqlxpkg

type Repository[T any] interface {
	Create(e *T) error

	ReadMany(limit, offset int) ([]T, error)

	ReadOne(id string) (T, error)

	Update(id string, e *T) (int64, error)

	Delete(id string) (int64, error)
}
