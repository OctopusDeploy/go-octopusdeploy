package service

import (
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/stretchr/testify/require"
)

type TestQuery struct {
	AccountType string   `uri:"accountType"`
	ID          string   `uri:"id"`
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

func TestFormatter(t *testing.T) {
	accountsURI := "/api/Spaces-1/accounts{/id}{?skip,take,ids,partialName,accountType}"
	url, err := url.ParseRequestURI(accountsURI)
	require.NoError(t, err)
	require.NotNil(t, url)

	template, err := uritemplates.Parse(accountsURI)
	require.NoError(t, err)

	templateString, err := template.Expand(make(map[string]interface{}))
	require.NoError(t, err)
	require.NotNil(t, templateString)

	rawURI := template.Names()
	require.NoError(t, err)
	require.NotNil(t, rawURI)

	test := TestQuery{}
	require.NotNil(t, test)

	test = TestQuery{
		AccountType: "UsernamePassword",
		ID:          "foo",
		IDs:         []string{"Foo", "Bar"},
		PartialName: "Foo Bar",
		Skip:        0,
		Take:        20,
	}

	values := make(map[string]interface{})
	values["accountType"] = test.AccountType
	values["id"] = test.ID
	values["ids"] = "Foo,Bar"
	values["partialName"] = test.PartialName
	values["skip"] = test.Skip
	values["take"] = test.Take

	expected, err := template.Expand(values)
	require.NoError(t, err)
	require.NotNil(t, expected)

	actual, err := template.Expand(test)
	require.NoError(t, err)
	require.NotNil(t, actual)

	require.Equal(t, expected, actual)
}
