package credentials_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestResourceWithUsernamePasswordAsJSON(t *testing.T) {
	description := internal.GetRandomName()
	id := internal.GetRandomName()
	name := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	selfLink := internal.GetRandomName()
	username := internal.GetRandomName()

	usernamePassword := credentials.NewUsernamePassword(username, password)

	usernamePasswordAsJSON, err := json.Marshal(usernamePassword)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordAsJSON)

	resource := credentials.NewResource(name, usernamePassword)
	resource.Description = description
	resource.ID = id
	resource.Links["Self"] = selfLink

	expectedJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Details": %s,
		"Id": "%s",
		"Name": "%s",
		"Links": {
			"Self": "%s"
		}
	}`, description, usernamePasswordAsJSON, id, name, selfLink)

	resourceAsJSON, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, resourceAsJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(resourceAsJSON))
}

func TestResourceWithAnonymousAsJSON(t *testing.T) {
	anonymous := credentials.NewAnonymous()
	description := internal.GetRandomName()
	id := internal.GetRandomName()
	name := internal.GetRandomName()
	selfLink := internal.GetRandomName()

	anonymousdAsJSON, err := json.Marshal(anonymous)
	require.NoError(t, err)
	require.NotNil(t, anonymousdAsJSON)

	resource := credentials.NewResource(name, anonymous)
	resource.Description = description
	resource.ID = id
	resource.Links["Self"] = selfLink

	expectedJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Details": %s,
		"Id": "%s",
		"Name": "%s",
		"Links": {
			"Self": "%s"
		}
	}`, description, anonymousdAsJSON, id, name, selfLink)

	resourceAsJSON, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, resourceAsJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(resourceAsJSON))
}

func TestResourceWithReferenceAsJSON(t *testing.T) {
	description := internal.GetRandomName()
	id := internal.GetRandomName()
	name := internal.GetRandomName()
	selfLink := internal.GetRandomName()

	reference := credentials.NewReference(id)

	referenceAsJSON, err := json.Marshal(reference)
	require.NoError(t, err)
	require.NotNil(t, referenceAsJSON)

	resource := credentials.NewResource(name, reference)
	resource.Description = description
	resource.ID = id
	resource.Links["Self"] = selfLink

	expectedJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Details": %s,
		"Id": "%s",
		"Name": "%s",
		"Links": {
			"Self": "%s"
		}
	}`, description, referenceAsJSON, id, name, selfLink)

	resourceAsJSON, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, resourceAsJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(resourceAsJSON))
}
