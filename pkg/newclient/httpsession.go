package newclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HttpSession is a layer over http.Client, and provides the following additional functionality:
// - Holding a 'base' URL, and using it to qualify relative URL's
// - Holding default HTTP request headers, and propagating them onto
// - Converting payloads to/from JSON, including our behaviour of converting HTTP 4xx/5xx responses into go error structs
// - Helpers for convenient Get/Put/Post/etc
// - Dealing with quirks of the go io subsystem (flushing buffers where required etc).
//
// While this borrows some ideas and quirk-handling from Sling, we don't want to use Sling itself because it has a weird API,
// and doesn't allow us access to some things that we need.
type HttpSession struct {
	HttpClient     *http.Client
	BaseURL        *url.URL
	DefaultHeaders map[string]string
}

// DoRawRequest adds any default headers to the HTTP request, and resolve the URL against the baseURL, then performs the HTTP request.
// This is the bottom of the stack of abstraction layers; for most situations a higher level method will be preferable
func (h *HttpSession) DoRawRequest(req *http.Request) (*http.Response, error) {
	if h.HttpClient == nil {
		return nil, errors.New("HttpSession.HttpClient is nil, can't DoRequest")
	}

	for k, v := range h.DefaultHeaders {
		if req.Header.Get(k) == "" {
			// only set DefaultHeaders if the request did not already set them
			req.Header.Set(k, v)
		}
	}

	if h.BaseURL != nil {
		// a HTTP URL with a path needs to retain that path during ResolveReference by removing any leading / from the request URL
		if h.BaseURL.Path != "" && strings.HasSuffix(h.BaseURL.Path, "/") {
			req.URL, _ = url.Parse(strings.TrimLeft(req.URL.String(), "/"))
		}
		req.URL = h.BaseURL.ResolveReference(req.URL)
	}

	return h.HttpClient.Do(req)
}

// DoRawJsonRequest layers JSON serialization over DoRawRequest
// outputResponseBody and outputResponseError should be pointers to structs which will receive unmarshaled JSON
// For most situations a higher level generic method like Get or Post will be better
func (h *HttpSession) DoRawJsonRequest(req *http.Request, requestBody any, outputResponseBody any, outputResponseError error) (*http.Response, error) {
	if req.Header == nil {
		req.Header = make(http.Header, 2)
	}

	if requestBody != nil {
		if req.Body != nil {
			return nil, errors.New("can't supply a requestBody to DoJsonRequest where the http.Request already has a body set")
		}

		body, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Body = io.NopCloser(bytes.NewReader(body))
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json,text/*;q=0.99")
	}

	resp, err := h.DoRawRequest(req)
	if err != nil {
		return nil, err
	}
	defer CloseResponse(resp)

	if resp.StatusCode == http.StatusNoContent || resp.ContentLength == 0 {
		// TODO the ContentLength check is copied from Sling, but it's valid for servers to stream responses
		// without a known content length. This won't handle such responses, which would be a bug. The octopus server tends not
		// to do this, so we can defer a fix until it becomes neccessary
		return resp, nil
	}

	bodyDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err = bodyDecoder.Decode(outputResponseBody)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else {
		// don't use core.APIErrorChecker, it's overly helpful and gets in the way of error handling.
		err = bodyDecoder.Decode(outputResponseError)
		if err != nil {
			return nil, err
		}
		if outputRes, ok := outputResponseError.(*core.APIError); ok {
			outputRes.StatusCode = resp.StatusCode
			return nil, outputRes
		}
		return nil, outputResponseError
	}
}

func DoRequest[TResponse any](httpSession *HttpSession, method string, path string, body any) (*TResponse, error) {
	pathUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: method,
		URL:    pathUrl,
		// body will be assigned inside DoRawJsonRequest
	}

	var responsePayload = new(TResponse)
	var errorPayload = new(core.APIError)
	_, err = httpSession.DoRawJsonRequest(req, body, responsePayload, errorPayload)
	if err != nil {
		return nil, err
	}

	// errorPayload is never nil because we allocated it just above
	if errorPayload.StatusCode != 0 {
		return nil, errorPayload
	}
	return responsePayload, nil
}

func DoDelete(httpSession *HttpSession, method string, path string) error {
	pathUrl, err := url.Parse(path)
	if err != nil {
		return err
	}

	req := &http.Request{
		Method: method,
		URL:    pathUrl,
	}

	var responsePayload = new(any)
	var errorPayload = new(core.APIError)
	_, err = httpSession.DoRawJsonRequest(req, nil, responsePayload, errorPayload)
	if err != nil {
		return err
	}

	// errorPayload is never nil because we allocated it just above
	if errorPayload.StatusCode != 0 {
		return errorPayload
	}
	return nil
}

func Get[TResponse any](httpSession *HttpSession, url string) (*TResponse, error) {
	return DoRequest[TResponse](httpSession, http.MethodGet, url, nil)
}

func Post[TResponse any](httpSession *HttpSession, url string, body any) (*TResponse, error) {
	return DoRequest[TResponse](httpSession, http.MethodPost, url, body)
}

func Put[TResponse any](httpSession *HttpSession, url string, body any) (*TResponse, error) {
	return DoRequest[TResponse](httpSession, http.MethodPut, url, body)
}

func Delete(httpSession *HttpSession, url string) error {
	return DoDelete(httpSession, http.MethodDelete, url)
}

// CloseResponse closes a response body; If you use DoRequest, and not one of the higher level helpers like Get or Post,
// you should use `defer CloseResponse(response)` to ensure it gets cleaned up properly
func CloseResponse(response *http.Response) {
	// behaviour copied from Sling

	// The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	// See: https://golang.org/pkg/net/http/#Response
	_, _ = io.Copy(io.Discard, response.Body)

	// when err is nil, resp contains a non-nil resp.Body which must be closed
	_ = response.Body.Close()
}
