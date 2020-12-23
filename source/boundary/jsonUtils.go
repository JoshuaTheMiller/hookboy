package boundary

import (
	"encoding/json"
)

// SerializeToNiceJSON transforms an object into an opinionatedly pretty string
func SerializeToNiceJSON(obj interface{}) (string, error) {
	b, err := json.MarshalIndent(obj, "", "	")

	if err != nil {
		return "", err
	}

	return string(b), nil
}
