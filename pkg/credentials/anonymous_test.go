package credentials_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestAnonymousMarshalJSON(t *testing.T) {
	anonymous := credentials.NewAnonymous()

	anonymousAsJSON, err := json.Marshal(anonymous)
	require.NoError(t, err)
	require.NotNil(t, anonymousAsJSON)

	expectedJSON := fmt.Sprintf(`{
		"Type": "%s"
	}`, credentials.GitCredentialTypeAnonymous)

	jsonassert.New(t).Assertf(expectedJSON, string(anonymousAsJSON))
}
