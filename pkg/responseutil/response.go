package responseutil

type Response[T any] struct {
	Data any `json:"data"`
}
