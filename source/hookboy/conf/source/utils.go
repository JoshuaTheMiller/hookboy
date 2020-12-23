package source

import (
	"fmt"
	"io/ioutil"
	"os"
)

type fileSystemObjectCheckResult string

const (
	doesNotExist fileSystemObjectCheckResult = "ObjectDoesNotExist"
	file         fileSystemObjectCheckResult = "IsAFile"
	folder       fileSystemObjectCheckResult = "IsAFolder"
)

func checkForFileSystemObject(path string) fileSystemObjectCheckResult {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return doesNotExist
	}

	if fileInfo.IsDir() {
		return folder
	}

	return file
}

func readFile(path string) ([]byte, error) {
	var result = checkForFileSystemObject(path)

	if result == doesNotExist {
		return nil, ConfigurationSourceError{
			Description: fmt.Sprintf("Cannot read nonexistant file: %s", path),
		}
	}

	if result == folder {
		return nil, ConfigurationSourceError{
			Description: fmt.Sprintf("Cannot read '%s' as it is a folder, not a file.", path),
		}
	}

	if result == file {
		bytes, err := ioutil.ReadFile(path)

		if err != nil {
			return nil, ConfigurationSourceError{
				Description: fmt.Sprintf("Problem reading configuration file: %s", path),
			}
		}

		return bytes, nil
	}

	return nil, ConfigurationSourceError{
		Description: "Unsupported file reading result type",
	}
}
