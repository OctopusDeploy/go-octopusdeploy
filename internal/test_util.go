package internal

import (
	"fmt"

	uuid "github.com/google/uuid"
)

func GetRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:44] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}
