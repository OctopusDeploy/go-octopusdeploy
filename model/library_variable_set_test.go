package model

// func TestLibraryVariableSetGet(t *testing.T) {
// 	client := getFakeOctopusClient(t, "/api/libraryVariableSets/LibraryVariableSets-41", http.StatusOK, getLibraryVariableSetResponseJSON)
// 	libraryVariableSet, err := client.LibraryVariableSet.Get("LibraryVariableSets-41")

// 	assert.Nil(t, err)
// 	assert.Equal(t, "MySet", libraryVariableSet.Name)
// 	assert.Equal(t, "The Description", libraryVariableSet.Description)
// 	assert.Equal(t, "variableset-LibraryVariableSets-41", libraryVariableSet.VariableSetID)
// 	assert.Equal(t, enums.Variables, libraryVariableSet.ContentType)
// }

// const getLibraryVariableSetResponseJSON = `
// {
//   "Id": "LibraryVariableSets-41",
//   "Name": "MySet",
//   "Description": "The Description",
//   "VariableSetId": "variableset-LibraryVariableSets-41",
//   "ContentType": "Variables",
//   "Templates": [],
//   "Links": {
//     "Self": "/api/libraryvariablesets/LibraryVariableSets-481",
//     "Variables": "/api/variables/variableset-LibraryVariableSets-481"
//   }
// }`

// func TestValidateLibraryVariableSetValuesJustANamePasses(t *testing.T) {

// 	libraryVariableSet := NewLibraryVariableSet("My Set")

// 	assert.Nil(t, ValidateLibraryVariableSetValues(libraryVariableSet))
// }

// func TestValidateLibraryVariableSetValuesMissingNameFails(t *testing.T) {

// 	libraryVariableSet := &LibraryVariableSet{}

// 	assert.Error(t, ValidateLibraryVariableSetValues(libraryVariableSet))
// }
