package response

type Response[T any] struct {
	Data any `json:"data"`
}
