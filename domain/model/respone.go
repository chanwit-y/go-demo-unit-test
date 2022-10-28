package model

type Respone[T any] struct {
	Errors  any    `json:"errors"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
