package newclient

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestDoRawJsonRequestSetsGetBody(t *testing.T) {
	var sentRequest *http.Request
	httpSession := &HttpSession{
		HttpClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				sentRequest = req
				return &http.Response{
					StatusCode: http.StatusNoContent,
					Body:       io.NopCloser(strings.NewReader("")),
				}, nil
			}),
		},
	}

	requestURL, err := url.Parse("https://octopus.example/api/projects")
	require.NoError(t, err)
	req := &http.Request{Method: http.MethodPost, URL: requestURL}

	_, err = httpSession.DoRawJsonRequest(req, map[string]string{"Name": "test"}, nil, nil)
	require.NoError(t, err)

	require.NotNil(t, sentRequest.GetBody)
	for range 2 {
		body, err := sentRequest.GetBody()
		require.NoError(t, err)
		content, err := io.ReadAll(body)
		require.NoError(t, err)
		require.JSONEq(t, `{"Name":"test"}`, string(content))
	}
}
