package sqlx

type Repository[T any] interface {
	Create(e *T) error
	ReadMany() ([]T, error)
	ReadOne(id string) *T
	Update(id string, e *T) (int64, error)
	Delete(id string) (int64, error)
}
