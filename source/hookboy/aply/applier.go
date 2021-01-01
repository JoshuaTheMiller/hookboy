package aply

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

type applierboy struct {
	WriteFile func(filename string, content string) error
}

func writeFile(filename string, content string) error {
	var dir = path.Dir(filename)

	var error = os.MkdirAll(dir, os.ModePerm)

	if error != nil {
		return internal.AplyError{
			Description:   "Problem creating new folder for file",
			InternalError: error,
		}
	}

	return ioutil.WriteFile(filename, []byte(content), os.ModePerm)
}

func (ab *applierboy) instantiate() {
	ab.WriteFile = writeFile
}

// Install installs the hooks with the given configuration
func (ab applierboy) Install(configuration conf.Configuration, ftc []internal.FileToCreate) (string, error) {
	ab.instantiate()

	return writeFiles(ab.WriteFile, ftc)
}

func writeFiles(writeFile func(filename string, content string) error, filesToCreate []internal.FileToCreate) (string, error) {
	for _, ftc := range filesToCreate {
		var content = ftc.Contents
		var fullFileName = ftc.Path
		var createHookFileError = writeFile(fullFileName, content)

		if createHookFileError != nil {
			return "", createHookFileError
		}
	}

	return hooksInstalledMessage, nil
}

var hooksInstalledMessage = "Hooks installed!"
