package internal

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	encoder.Encode(data)
	return buffer.String()
}
