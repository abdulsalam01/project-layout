package base

type Response[T any] struct {
	Data    T
	Message string `json:"message"`
	Success bool   `json:"success"`
}
