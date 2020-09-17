package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestActionTemplateServiceURITemplate = "action-templates-service"
)

func TestNewActionTemplateService(t *testing.T) {
	service := NewActionTemplateService(nil, "")
	assert.Nil(t, service)
	createActionTemplateService(t)
}

func TestActionTemplateServiceGetWithEmptyID(t *testing.T) {
	service := createActionTemplateService(t)

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestActionTemplateGetWithEmptyName(t *testing.T) {
	service := createActionTemplateService(t)

	resource, err := service.GetByName("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestActionTemplateServiceAddWithNilActionTemplate(t *testing.T) {
	service := createActionTemplateService(t)

	account, err := service.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestActionTemplateServiceAddWithInvalidActionTemplate(t *testing.T) {
	service := createActionTemplateService(t)

	account, err := service.Add(&model.ActionTemplate{})

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestActionTemplateServiceDeleteWithEmptyID(t *testing.T) {
	service := createActionTemplateService(t)

	err := service.Delete("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))

	err = service.Delete(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))
}

func TestActionTemplateServiceGetWithInvalidID(t *testing.T) {
	service := createActionTemplateService(t)

	item, err := service.Get("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, item)

	item, err = service.Get(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, item)
}

func createActionTemplateService(t *testing.T) *ActionTemplateService {
	service := NewActionTemplateService(&sling.Sling{}, TestActionTemplateServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestActionTemplateServiceURITemplate)
	assert.Equal(t, service.name, "ActionTemplateService")

	return service
}
