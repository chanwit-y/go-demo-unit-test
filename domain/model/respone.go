package model

type Respone[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Error   bool   `json:"error"`
}
