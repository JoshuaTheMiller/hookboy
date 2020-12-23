package cli

import (
	"encoding/json"
)

func serializeToNiceJSON(obj interface{}) (string, error) {
	b, err := json.MarshalIndent(obj, "", "	")

	if err != nil {
		return "", err
	}

	return string(b), nil
}
