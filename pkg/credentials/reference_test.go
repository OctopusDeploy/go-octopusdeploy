package credentials_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestReferenceMarshalJSON(t *testing.T) {
	id := internal.GetRandomName()
	reference := credentials.NewReference(id)

	referenceAsJSON, err := json.Marshal(reference)
	require.NoError(t, err)
	require.NotNil(t, referenceAsJSON)

	expectedJSON := fmt.Sprintf(`{
		"Id": "%s",
		"Type": "%s"
	}`, id, credentials.GitCredentialTypeReference)

	jsonassert.New(t).Assertf(expectedJSON, string(referenceAsJSON))
}
