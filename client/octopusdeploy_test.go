package client

import (
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewClientWithEmptyProperties(t *testing.T) {
	client, err := NewClient(nil, "", "")

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientWithEmptyClient(t *testing.T) {
	client, err := NewClient(nil, "fake-url", "fake-api-key")

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientWithEmptyOctopusURL(t *testing.T) {
	client, err := NewClient(&http.Client{}, "", "fake-api-key")

	assert.Error(t, err)
	assert.Nil(t, client)

	client, err = NewClient(&http.Client{}, " ", "fake-api-key")

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientWithEmptyAPIKey(t *testing.T) {
	client, err := NewClient(&http.Client{}, "fake-url", "")

	assert.Error(t, err)
	assert.Nil(t, client)

	client, err = NewClient(&http.Client{}, "fake-url", " ")

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestForSpaceWithEmptyProperties(t *testing.T) {
	client, err := ForSpace(nil, "", "", nil)

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestForSpaceWithEmptyClient(t *testing.T) {
	client, err := ForSpace(nil, "fake-url", "fake-api-key", &model.Space{})

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestForSpaceWithEmptyOctopusURL(t *testing.T) {
	client, err := ForSpace(&http.Client{}, "", "fake-api-key", &model.Space{})

	assert.Error(t, err)
	assert.Nil(t, client)

	client, err = ForSpace(&http.Client{}, " ", "fake-api-key", &model.Space{})

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestForSpaceWithEmptyAPIKey(t *testing.T) {
	client, err := ForSpace(&http.Client{}, "fake-url", "", &model.Space{})

	assert.Error(t, err)
	assert.Nil(t, client)

	client, err = ForSpace(&http.Client{}, "fake-url", " ", &model.Space{})

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestForSpaceWithEmptySpace(t *testing.T) {
	client, err := ForSpace(&http.Client{}, "fake-url", "fake-api-key", nil)

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestGetWithEmptyParameters(t *testing.T) {
	resource, err := apiGet(nil, nil, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestGetWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	resource, err := apiGet(nil, input, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestGetWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	resource, err := apiGet(sling.New(), input, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = apiGet(sling.New(), input, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAddWithEmptyParameters(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiAdd(nil, input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAddWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiAdd(nil, input, response, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAddWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiAdd(sling.New(), input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = apiAdd(sling.New(), input, response, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestPostWithEmptyParameters(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiPost(nil, input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestPostWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiPost(nil, input, response, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestPostWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiPost(sling.New(), input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = apiPost(sling.New(), input, response, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestUpdateWithEmptyParameters(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiUpdate(nil, input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestUpdateWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiUpdate(nil, input, response, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestUpdateWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := apiUpdate(sling.New(), input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = apiUpdate(sling.New(), input, response, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestDeleteWithEmptyParameters(t *testing.T) {
	err := apiDelete(nil, "")
	assert.Error(t, err)
}

func TestDeleteWithEmptySling(t *testing.T) {
	err := apiDelete(nil, "fake-path")
	assert.Error(t, err)
}

func TestDeleteWithEmptyPath(t *testing.T) {
	err := apiDelete(nil, "")
	assert.Error(t, err)

	err = apiDelete(nil, " ")
	assert.Error(t, err)
}

type inputTestValueStruct struct {
	test string
}

func (i *inputTestValueStruct) GetID() string {
	return "fake-ID"
}

func (i *inputTestValueStruct) Validate() error {
	return nil
}

type inputTestResponseStruct struct {
	test string
}

func (i *inputTestResponseStruct) GetID() string {
	return "fake-ID"
}

func (i *inputTestResponseStruct) Validate() error {
	return nil
}
