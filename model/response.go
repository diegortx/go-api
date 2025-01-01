package model

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}
