package ephemeralenvironments

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) newclient.Client {
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	baseURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	return newclient.NewClientS(&newclient.HttpSession{
		HttpClient:     server.Client(),
		BaseURL:        baseURL,
		DefaultHeaders: map[string]string{},
	}, "Spaces-1")
}

func TestGetStatus(t *testing.T) {
	var method, requested string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		requested = r.URL.RequestURI()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"ProvisioningFailed"}`))
	})

	status, err := GetStatus(client, "Spaces-1", "Environments-1001", "Projects-1")
	require.NoError(t, err)

	assert.Equal(t, http.MethodGet, method)
	assert.Equal(t, "/api/Spaces-1/projects/Projects-1/environments/ephemeral/Environments-1001/status", requested)
	assert.Equal(t, "ProvisioningFailed", status.Status)
}

func TestGetStatusReturnsServerError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"ErrorMessage":"not found"}`))
	})

	status, err := GetStatus(client, "Spaces-1", "Environments-1001", "Projects-1")
	assert.Error(t, err)
	assert.Nil(t, status)
}

func TestGetStatusRequiresIDs(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("no request expected, got %s %s", r.Method, r.URL.RequestURI())
	})

	_, err := GetStatus(client, "Spaces-1", "", "Projects-1")
	assert.Error(t, err)

	_, err = GetStatus(client, "Spaces-1", "Environments-1001", "")
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	var method, requested string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		requested = r.URL.RequestURI()
		w.WriteHeader(http.StatusOK)
	})

	err := Delete(client, "Spaces-1", "Environments-1001")
	require.NoError(t, err)

	assert.Equal(t, http.MethodDelete, method)
	assert.Equal(t, "/api/Spaces-1/environments/Environments-1001", requested)
}

func TestDeleteRequiresID(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("no request expected, got %s %s", r.Method, r.URL.RequestURI())
	})

	assert.Error(t, Delete(client, "Spaces-1", ""))
}
