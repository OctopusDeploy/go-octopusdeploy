package api

import "os"

type Environment interface {
	Getenv(key string) string
}

type OsEnvironment struct {
}

type MockEnvironment struct {
	values map[string]string
}

func NewMockEnvironment() *MockEnvironment {
	return &MockEnvironment{
		values: make(map[string]string),
	}
}

func (e MockEnvironment) Getenv(key string) string {
	return e.values[key]
}

func (e MockEnvironment) Setenv(key string, value string) {
	e.values[key] = value
}

func (e OsEnvironment) Getenv(key string) string {
	return os.Getenv(key)
}
