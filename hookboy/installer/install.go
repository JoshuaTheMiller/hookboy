package installer

import (
	"errors"

	"github.com/mitchellh/go-homedir"
)

// GetHomeDir as
func GetHomeDir() (string, error) {
	dir, err := homedir.Dir()

	if err != nil {
		return "", errors.New("Unable to retrieve home directory")
	}

	return dir, nil
}
