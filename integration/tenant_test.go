package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	octopusClient = initTest()
}

func TestTenantAddAndDelete(t *testing.T) {
	tenantName := getRandomName()
	expected := getTestTenant(tenantName)
	actual := createTestTenant(t, tenantName)

	defer cleanTenant(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "tenant name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "tenant doesn't contain an ID from the octopus server")
}

func TestTenantAddGetAndDelete(t *testing.T) {
	tenant := createTestTenant(t, getRandomName())
	defer cleanTenant(t, tenant.ID)

	getTenant, err := octopusClient.Tenants.Get(tenant.ID)
	assert.Nil(t, err, "there was an error raised getting tenant when there should not be")
	assert.Equal(t, tenant.Name, getTenant.Name)
}

func TestTenantGetThatDoesNotExist(t *testing.T) {
	tenantID := "there-is-no-way-this-tenant-id-exists-i-hope"
	expected := client.ErrItemNotFound
	tenant, err := octopusClient.Tenants.Get(tenantID)

	assert.Error(t, err, "there should have been an error raised as this tenant should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, tenant, "no tenant should have been returned")
}

func TestTenantGetAll(t *testing.T) {
	// create many tenant to test pagination
	tenantsToCreate := 32
	sum := 0
	for i := 0; i < tenantsToCreate; i++ {
		tenant := createTestTenant(t, getRandomName())
		defer cleanTenant(t, tenant.ID)
		sum += i
	}

	allTenants, err := octopusClient.Tenants.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all tenants failed when it shouldn't: %s", err)
	}

	numberOfTenants := len(*allTenants)

	// check there are greater than or equal to the amount of tenants requested to be created, otherwise pagination isn't working
	if numberOfTenants < tenantsToCreate {
		t.Fatalf("There should be at least %d tenants created but there was only %d. Pagination is likely not working.", tenantsToCreate, numberOfTenants)
	}

	additionalTenant := createTestTenant(t, getRandomName())
	defer cleanTenant(t, additionalTenant.ID)

	allTenantsAfterCreatingAdditional, err := octopusClient.Tenants.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all tenants failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for tenant when not expected")
	assert.Equal(t, len(*allTenantsAfterCreatingAdditional), numberOfTenants+1, "created an additional tenant and expected number of tenants to increase by 1")
}

func TestTenantUpdate(t *testing.T) {
	tenant := createTestTenant(t, getRandomName())
	defer cleanTenant(t, tenant.ID)

	newTenantName := getRandomName()
	const newDescription = "this should be updated"
	// const newSkipMachineBehavior = "SkipUnavailableMachines"

	tenant.Name = newTenantName
	tenant.Description = newDescription

	updatedTenant, err := octopusClient.Tenants.Update(&tenant)
	assert.Nil(t, err, "error when updating tenant")
	assert.Equal(t, newTenantName, updatedTenant.Name, "tenant name was not updated")
	assert.Equal(t, newDescription, updatedTenant.Description, "tenant description was not updated")
}

func TestTenantGetByName(t *testing.T) {
	tenant := createTestTenant(t, getRandomName())
	defer cleanTenant(t, tenant.ID)

	foundTenant, err := octopusClient.Tenants.GetByName(tenant.Name)
	assert.Nil(t, err, "error when looking for tenant when not expected")
	assert.Equal(t, tenant.Name, foundTenant.Name, "tenant not found when searching by its name")
}

func createTestTenant(t *testing.T, tenantName string) model.Tenant {
	p := getTestTenant(tenantName)
	createdTenant, err := octopusClient.Tenants.Add(&p)

	if err != nil {
		t.Fatalf("creating tenant %s failed when it shouldn't: %s", tenantName, err)
	}

	return *createdTenant
}

func getTestTenant(tenantName string) model.Tenant {
	p := model.NewTenant(tenantName, "Lifecycles-1")

	return *p
}

func cleanTenant(t *testing.T, tenantID string) {
	err := octopusClient.Tenants.Delete(tenantID)

	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting tenant failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
