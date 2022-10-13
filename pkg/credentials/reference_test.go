package credentials

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestReferenceMarshalJSON(t *testing.T) {
	id := internal.GetRandomName()
	reference := NewReference(id)

	referenceAsJSON, err := json.Marshal(reference)
	require.NoError(t, err)
	require.NotNil(t, referenceAsJSON)

	expectedJSON := fmt.Sprintf(`{
		"Id": "%s",
		"Type": "%s"
	}`, id, GitCredentialTypeReference)

	jsonassert.New(t).Assertf(expectedJSON, string(referenceAsJSON))
}
