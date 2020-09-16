package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewActionTemplateServiceWithNil(t *testing.T) {
	service := NewActionTemplateService(nil, "")

	assert.Nil(t, service)
}

func TestNewActionTemplateServiceWithEmptyClient(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{}, "actiontemplates")

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")
	assert.NotNil(t, service.sling)
}

func TestActionTemplateServiceGetWithEmptyID(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{}, "actiontemplates")

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestActionTemplateGetWithEmptyName(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{}, "actiontemplates")

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")

	resource, err := service.GetByName("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestActionTemplateServiceAddWithNilActionTemplate(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{}, "actiontemplates")

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")

	account, err := service.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestActionTemplateServiceAddWithInvalidActionTemplate(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{}, "actiontemplates")

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")

	account, err := service.Add(&model.ActionTemplate{})

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestActionTemplateServiceDeleteWithEmptyID(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{}, "actiontemplates")

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")

	err := service.Delete("")

	assert.Error(t, err)

	err = service.Delete(" ")

	assert.Error(t, err)
}
