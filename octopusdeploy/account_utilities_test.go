package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToAccount(t *testing.T) {
	account, err := ToAccount(nil)
	require.Nil(t, account)
	require.Error(t, err)
}
