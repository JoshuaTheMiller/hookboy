package boundary

import (
	"errors"

	"github.com/mitchellh/go-homedir"
)

// GetHomeDir to retrieve the home directory as configured by the current
// system
func GetHomeDir() (string, error) {
	dir, err := homedir.Dir()

	if err != nil {
		return "", errors.New("Unable to retrieve home directory")
	}

	return dir, nil
}
