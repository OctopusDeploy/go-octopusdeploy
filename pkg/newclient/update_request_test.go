package newclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeBase struct {
	ID    string            `json:"Id,omitempty"`
	Links map[string]string `json:"Links,omitempty"`
}

type fakeChannel struct {
	Name                   string            `json:"Name,omitempty"`
	Description            string            `json:"Description,omitempty"`
	IsDefault              bool              `json:"IsDefault"`
	CustomFieldDefinitions []fakeCustomField `json:"CustomFieldDefinitions,omitempty"`
	Tags                   []string          `json:"Tags,omitempty"`
	fakeBase
}

type fakeCustomField struct {
	FieldName string `json:"FieldName"`
}

func TestUpdateRequest_DefaultMatchesPlainMarshal(t *testing.T) {
	c := &fakeChannel{Name: "n", fakeBase: fakeBase{ID: "Channels-1"}}

	got, err := json.Marshal(NewUpdateRequest(c))
	require.NoError(t, err)
	want, err := json.Marshal(c)
	require.NoError(t, err)
	require.JSONEq(t, string(want), string(got))
}

func TestUpdateRequest_ClearEmitsEmptySliceWhenNil(t *testing.T) {
	c := &fakeChannel{Name: "n"}

	got, err := json.Marshal(NewUpdateRequest(c).Clear("CustomFieldDefinitions"))
	require.NoError(t, err)
	require.Contains(t, string(got), `"CustomFieldDefinitions":[]`)
}

func TestUpdateRequest_ClearOverwritesNonEmptySlice(t *testing.T) {
	c := &fakeChannel{
		Name:                   "n",
		CustomFieldDefinitions: []fakeCustomField{{FieldName: "x"}},
	}

	got, err := json.Marshal(NewUpdateRequest(c).Clear("CustomFieldDefinitions"))
	require.NoError(t, err)
	require.Contains(t, string(got), `"CustomFieldDefinitions":[]`)
	require.NotContains(t, string(got), `"FieldName":"x"`)
}

func TestUpdateRequest_ClearEmitsZeroForPrimitive(t *testing.T) {
	c := &fakeChannel{Name: "n", Description: "still here"}

	got, err := json.Marshal(NewUpdateRequest(c).Clear("Description"))
	require.NoError(t, err)
	require.Contains(t, string(got), `"Description":""`)
}

func TestUpdateRequest_ClearWorksOnEmbeddedField(t *testing.T) {
	c := &fakeChannel{Name: "n"}

	got, err := json.Marshal(NewUpdateRequest(c).Clear("Id"))
	require.NoError(t, err)
	require.Contains(t, string(got), `"Id":""`)
}

func TestUpdateRequest_ClearUnknownFieldErrorsAtMarshal(t *testing.T) {
	c := &fakeChannel{Name: "n"}

	_, err := json.Marshal(NewUpdateRequest(c).Clear("NotAField"))
	require.Error(t, err)
	require.Contains(t, err.Error(), "NotAField")
}

func TestUpdateRequest_MultipleClearsCompose(t *testing.T) {
	c := &fakeChannel{Name: "n", Description: "x"}

	got, err := json.Marshal(NewUpdateRequest(c).
		Clear("CustomFieldDefinitions").
		Clear("Tags").
		Clear("Description"))
	require.NoError(t, err)
	require.Contains(t, string(got), `"CustomFieldDefinitions":[]`)
	require.Contains(t, string(got), `"Tags":[]`)
	require.Contains(t, string(got), `"Description":""`)
}

func TestUpdateRequest_ResourceReturnsWrappedPointer(t *testing.T) {
	c := &fakeChannel{Name: "n"}
	req := NewUpdateRequest(c)
	require.Same(t, c, req.Resource())
}
