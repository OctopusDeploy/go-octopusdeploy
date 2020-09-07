package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type LibraryVariableSetUsageEntry struct {
	LibraryVariableSetID   string `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`
	Resource
}

func (l *LibraryVariableSetUsageEntry) GetID() string {
	return l.ID
}

// Validate returns a collection of validation errors against the library
// variable set usage entry's internal values.
func (l *LibraryVariableSetUsageEntry) Validate() error {
	validate := validator.New()
	err := validate.Struct(l)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}

		return err
	}

	return nil
}

var _ ResourceInterface = &LibraryVariableSetUsageEntry{}
