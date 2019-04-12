package octopusdeploy

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLibraryVariableSetGet(t *testing.T) {

	httpClient := http.Client{}
	httpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "/api/libraryVariableSets/LibraryVariableSets-41", r.URL.Path)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(getLibraryVariableSetResponseJSON)),
		}, nil
	})

	client := getFakeOctopusClient(httpClient)
	libraryVariableSet, err := client.LibraryVariableSet.Get("LibraryVariableSets-41")

	assert.Nil(t, err)
	assert.Equal(t, "MySet", libraryVariableSet.Name)
	assert.Equal(t, "The Description", libraryVariableSet.Description)
	assert.Equal(t, "variableset-LibraryVariableSets-41", libraryVariableSet.VariableSetId)
	assert.Equal(t, VariableSetContentType_Variables, libraryVariableSet.ContentType)
}

const getLibraryVariableSetResponseJSON = `
{
  "Id": "LibraryVariableSets-41",
  "Name": "MySet",
  "Description": "The Description",
  "VariableSetId": "variableset-LibraryVariableSets-41",
  "ContentType": "Variables",
  "Templates": [],
  "Links": {
    "Self": "/api/libraryvariablesets/LibraryVariableSets-481",
    "Variables": "/api/variables/variableset-LibraryVariableSets-481"
  }
}`

func TestValidateLibraryVariableSetValuesJustANamePasses(t *testing.T) {

	libraryVariableSet := NewLibraryVariableSet("My Set")

	assert.Nil(t, ValidateLibraryVariableSetValues(libraryVariableSet))
}

func TestValidateLibraryVariableSetValuesMissingNameFails(t *testing.T) {

	libraryVariableSet := &LibraryVariableSet{}

	assert.Error(t, ValidateLibraryVariableSetValues(libraryVariableSet))
}
