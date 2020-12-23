package source

import "os"

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
