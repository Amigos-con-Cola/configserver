package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type optype string

const (
	getAll optype = "GET_ALL"
	getOne optype = "GET_ONE"
	set    optype = "SET"
)

var (
	apiBaseUrl string = "http://localhost:3000"
)

type operation interface {
	getUrl() string
	getMethod() string
	getBody() map[string]string
}

// -- GET OP

type getOp struct {
	env string
	key string
}

func NewGet(env string, key string) operation {
	return &getOp{
		env: env,
		key: key,
	}
}

func (gop *getOp) getUrl() string {
	return "api/v1/" + gop.env + "/" + gop.key
}

func (gop *getOp) getMethod() string {
	return http.MethodGet
}

func (gop *getOp) getBody() map[string]string {
	return nil
}

// -- GET ALL OP

type getAllOp struct {
	env string
}

func NewGetAll(env string) operation {
	return &getAllOp{env: env}
}

func (ga *getAllOp) getUrl() string {
	return "/api/v1/" + ga.env
}

func (ga *getAllOp) getMethod() string {
	return http.MethodGet
}

func (ga *getAllOp) getBody() map[string]string {
	return nil
}

// -- SET OP

type setOp struct {
	env   string
	key   string
	value string
}

func NewSet(env string, key string, value string) operation {
	return &setOp{
		env:   env,
		key:   key,
		value: value,
	}
}

func (sop *setOp) getUrl() string {
	return "/api/v1/" + sop.env
}

func (sop *setOp) getMethod() string {
	return http.MethodPost
}

func (sop *setOp) getBody() map[string]string {
	b := make(map[string]string)
	b[sop.key] = sop.value
	return b
}

// -- MAIN CLIENT API

func DoPerform(operation operation) (map[string]string, error) {
	apiUrl := os.Getenv("CONFIG_SERVER_BASE_URL")
	if apiUrl != "" {
		apiBaseUrl = apiUrl
	}

	var bodyBytes *bytes.Buffer
	opBody := operation.getBody()

	if opBody != nil {
		json, err := json.Marshal(opBody)
		if err != nil {
			return nil, err
		}
		bodyBytes = bytes.NewBuffer(json)
	}

	req, err := http.NewRequest(operation.getMethod(), apiUrl+operation.getUrl(), bodyBytes)
	if err != nil {
		return nil, err
	}

	username := os.Getenv("CONFIG_SERVER_USERNAME")
	password := os.Getenv("CONFIG_SERVER_PASSWORD")

	if username == "" || password == "" {
		return nil, errors.New("Credentials missing to authenticate against config server")
	}

	req.SetBasicAuth(username, password)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		var responseBody map[string]string
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, &responseBody)
		return responseBody, nil
	case http.StatusUnauthorized:
		return nil, errors.New("The request is unauthorized")
	case http.StatusBadRequest:
		return nil, errors.New("Malformed request")
	}

	return nil, errors.New("API returned unexpected status code")
}
