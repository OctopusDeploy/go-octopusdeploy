package resources

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

type NotAllTestStruct struct {
	Name string `validate:"notall"`
}

func TestNotAllValidation(t *testing.T) {
	notAll := &NotAllTestStruct{
		Name: "All",
	}

	v := validator.New()

	err := v.RegisterValidation("notall", validation.NotAll)
	require.NoError(t, err)

	err = v.Struct(notAll)
	require.Error(t, err)

	notAll.Name = "all"

	err = v.Struct(notAll)
	require.Error(t, err)

	notAll.Name = "ALL"

	err = v.Struct(notAll)
	require.Error(t, err)

	notAll.Name = "aLl"

	err = v.Struct(notAll)
	require.Error(t, err)
}
