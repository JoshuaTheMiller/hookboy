package source

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hookboy/source/hookboy"
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

func readFile(path string) ([]byte, hookboy.Error) {
	var result = checkForFileSystemObject(path)

	if result == doesNotExist {
		return nil, hookboy.NewError(fmt.Sprintf("Cannot read nonexistant file: %s", path))
	}

	if result == folder {
		return nil, hookboy.NewError(fmt.Sprintf("Cannot read '%s' as it is a folder, not a file.", path))
	}

	if result == file {
		bytes, err := ioutil.ReadFile(path)

		if err != nil {
			return nil, hookboy.NewError(fmt.Sprintf("Problem reading configuration file: %s", path))
		}

		return bytes, nil
	}

	return nil, hookboy.NewError("Unsupported file reading result type")
}
