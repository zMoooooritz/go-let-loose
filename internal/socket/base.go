package socket

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StatusCode int

const (
	StatusOk            StatusCode = 200
	StatusBadRequest    StatusCode = 400
	StatusUnauthorized  StatusCode = 401
	StatusInternalError StatusCode = 500
)

type RconError struct {
	Code    StatusCode
	Message string
}

func (e RconError) Error() string {
	return fmt.Sprintf("RconError: %d - %s", e.Code, e.Message)
}

func NewRconError(code StatusCode, message string) RconError {
	return RconError{
		Code:    code,
		Message: message,
	}
}

type RawRequest struct {
	AuthToken   string `json:"AuthToken"`
	Version     int    `json:"Version"`
	Name        string `json:"Name"`
	ContentBody string `json:"ContentBody"`
}

func NewRawRequest(authToken string, Version int, name string, contentBody string) RawRequest {
	return RawRequest{
		AuthToken:   authToken,
		Version:     Version,
		Name:        name,
		ContentBody: contentBody,
	}
}

func (r RawRequest) String() string {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}

func (r RawRequest) Pack() []byte {
	body, err := json.Marshal(r)
	if err != nil {
		return []byte{}
	}
	return body
}

type RconRequest[T any] struct {
	Body T
}

func (r *RconRequest[T]) ToArgs() (string, string) {
	body := r.Body
	var d []byte
	t := reflect.ValueOf(body)
	if t.Kind() == reflect.String {
		d = []byte(t.String())
	} else {
		d, _ = json.Marshal(body)
	}
	cmd := reflect.TypeOf(body).Name()
	return cmd, string(d)
}

type RconResponse struct {
	StatusCode    StatusCode `json:"StatusCode"`
	StatusMessage string     `json:"StatusMessage"`
	Version       int        `json:"Version"`
	Name          string     `json:"Name"`
	ContentBody   string     `json:"ContentBody"`
}

func (r RconResponse) String() string {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}
