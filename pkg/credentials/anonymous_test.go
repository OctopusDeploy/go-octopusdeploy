package credentials

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestAnonymousMarshalJSON(t *testing.T) {
	anonymous := NewAnonymous()

	anonymousAsJSON, err := json.Marshal(anonymous)
	require.NoError(t, err)
	require.NotNil(t, anonymousAsJSON)

	expectedJSON := fmt.Sprintf(`{
		"Type": "%s"
	}`, GitCredentialTypeAnonymous)

	jsonassert.New(t).Assertf(expectedJSON, string(anonymousAsJSON))
}
