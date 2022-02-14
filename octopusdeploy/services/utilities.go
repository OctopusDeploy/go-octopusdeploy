package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	empty = ""
	tab   = "\t"
)

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func isAPIKey(apiKey string) bool {
	if len(apiKey) < 5 {
		return false
	}

	var expression = regexp.MustCompile(`^(API-)([A-Z\d])+$`)
	return expression.MatchString(apiKey)
}

func PrettyJSON(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	encoder.Encode(data)
	return buffer.String()
}

func trimTemplate(uri string) string {
	return strings.Split(uri, "{")[0]
}

func createBuiltInTeamsCannotDeleteError() error {
	return fmt.Errorf("The built-in teams cannot be deleted.")
}

func createInvalidClientStateError(ServiceName string) error {
	return fmt.Errorf("%s: the state of the internal Client is invalid", ServiceName)
}

func createInvalidPathError(ServiceName string) error {
	return fmt.Errorf("%s: the internal path is not set", ServiceName)
}

func createItemNotFoundError(ServiceName string, methodName string, name string) error {
	return fmt.Errorf("%s: the item (%s) via %s was not found", ServiceName, name, methodName)
}

func createClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal Client", methodName)
}

func createResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func createValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }
