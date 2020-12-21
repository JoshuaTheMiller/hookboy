package installer

import (
	"errors"

	"github.com/mitchellh/go-homedir"
)

func getHomeDir() (string, error) {
	dir, err := homedir.Dir()

	if err != nil {
		return "", errors.New("Unable to retrieve home directory")
	}

	return dir, nil
}
