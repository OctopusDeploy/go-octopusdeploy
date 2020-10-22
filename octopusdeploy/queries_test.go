package octopusdeploy

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/require"
)

type TestQuery struct {
	AccountType string   `url:"accountType,omitempty"`
	IDs         []string `url:"ids,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
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

	test := TestQuery{
		AccountType: "UsernamePassword",
		IDs:         []string{"Foo", "Bar"},
		PartialName: "Foo Bar",
		Take:        20,
	}

	v, _ := query.Values(test)
	fmt.Print(v.Encode())
}
