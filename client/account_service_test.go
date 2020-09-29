package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestAccountService(t *testing.T) {
	t.Run("New", TestNewAccountService)
	t.Run("Parameters", TestAccountServiceParameters)
	t.Run("GetAll", TestAccountServiceGetAll)
	t.Run("GetByID", TestAccountServiceGetByID)
	t.Run("GetByName", TestAccountServiceGetByName)
	t.Run("GetByAccountType", TestAccountServiceGetByAccountType)
	t.Run("Add", TestAccountServiceAdd)
	t.Run("Update", TestAccountServiceUpdateWithEmptyAccount)
	t.Run("Usage", TestAccountServiceGetUsages)
}

func TestNewAccountService(t *testing.T) {
	serviceFunction := newAccountService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAccountService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *accountService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestAccountServiceGetByAccountType(t *testing.T) {
	assert := assert.New(t)

	service := createAccountService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	for _, typeName := range enum.AccountTypeNames() {
		accountType, err := enum.ParseAccountType(typeName)

		assert.NoError(err)
		if err != nil {
			return
		}

		resourceList, err := service.GetByAccountType(accountType)

		assert.NoError(err)

		if len(resourceList) > 0 {
			resourceToCompare, err := service.GetByID(resourceList[0].ID)

			assert.NoError(err)
			assert.EqualValues(resourceList[0], *resourceToCompare)
		}
	}
}

func TestAccountServiceGetByID(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)

	if len(resourceList) > 0 {
		resourceToCompare, err := service.GetByID(resourceList[0].ID)

		assert.NoError(err)
		assert.EqualValues(resourceList[0], *resourceToCompare)
	}

	value := getRandomName()
	resource, err := service.GetByID(value)

	assert.Equal(err, createResourceNotFoundError("account", "ID", value))
	assert.Nil(resource)
}

func TestAccountServiceGetByName(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)

	if len(resourceList) > 0 {
		resourceToCompare, err := service.GetByName(resourceList[0].Name)

		assert.NoError(err)
		assert.EqualValues(*resourceToCompare, resourceList[0])
	}
}

func TestAccountServiceGetAll(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)
}

func TestAccountServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := createAccountService(t)
			assert := assert.New(t)

			assert.NotNil(service)
			if service == nil {
				return
			}

			resource, err := service.GetByID(tc.parameter)

			assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(resource)

			resourceList, err := service.GetByPartialName(tc.parameter)

			assert.Equal(createInvalidParameterError(operationGetByPartialName, parameterName), err)
			assert.NotNil(resourceList)

			err = service.DeleteByID(tc.parameter)

			assert.Error(err)
			assert.Equal(err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func TestAccountServiceGetUsages(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	accounts, err := service.GetAll()

	assert.NoError(err)

	if len(accounts) > 0 {
		accountUsages, err := service.GetUsages(accounts[0])

		assert.NoError(err)
		assert.NotNil(accountUsages)
	}
}

func TestAccountServiceGetByIDs(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)

	idList := []string{}
	for _, resource := range resourceList {
		idList = append(idList, resource.ID)
	}

	resourceListToCompare, err := service.GetByIDs(idList)

	assert.NoError(err)
	assert.Equal(len(resourceList), len(resourceListToCompare))
}

func TestAccountServiceAdd(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	resource, err := service.Add(nil)

	assert.Equal(err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(resource)

	resource, err = service.Add(&model.Account{})

	assert.Error(err)
	assert.Nil(resource)

	resource, err = model.NewUsernamePasswordAccount(getRandomName())

	assert.NoError(err)
	assert.NotNil(resource)

	if err != nil {
		return
	}

	resource, err = service.Add(resource)

	assert.NoError(err)
	assert.NotNil(resource)

	err = service.DeleteByID(resource.ID)

	assert.NoError(err)
}

func TestAccountServiceUpdateWithEmptyAccount(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	account, err := service.Update(model.Account{})

	assert.Error(err)
	assert.Nil(account)
}

func TestAccountServiceUpdate(t *testing.T) {
	service := createAccountService(t)
	assert := assert.New(t)

	resource, err := model.NewUsernamePasswordAccount(getRandomName())

	assert.NoError(err)
	assert.NotNil(resource)

	if err != nil {
		return
	}

	resourceToCompare, err := service.Add(resource)

	assert.NoError(err)
	assert.NotNil(resourceToCompare)

	resourceToCompare.Name = getRandomName()

	updatedResource, err := service.Update(*resourceToCompare)

	assert.NoError(err)
	assert.Equal(resourceToCompare.Name, updatedResource.Name)
}

func createAccountService(t *testing.T) *accountService {
	service := newAccountService(nil, TestURIAccounts)
	testNewService(t, service, TestURIAccounts, serviceAccountService)
	return service
}
