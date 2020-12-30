package aply

import (
	"io/ioutil"
	"os"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

// GetApplier returns an object that implements the Applier interface
func GetApplier() Applier {
	return &applierboy{}
}

// Applier configures hooks
type Applier interface {
	Install(configuration conf.Configuration) (string, error)
}

type applierboy struct {
	WriteFile func(filename string, content string) error
}

func writeFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), os.ModePerm)
}

func (ab *applierboy) instantiate() {
	ab.WriteFile = writeFile
}

// Install installs the hooks with the given configuration
func (ab *applierboy) Install(configuration conf.Configuration) (string, error) {
	// TODO: should most likely accept files to create as a parameter, thus eliminating another
	// dependency here.
	// FileToCreate should be moved to an internal package as well (same with many of these interfaces)
	ab.instantiate()

	var prepboy = internal.GetPrepper()

	var filesToCreate, _ = prepboy.PrepareHookfileInfo(configuration)

	return writeFiles(ab.WriteFile, filesToCreate)
}

func writeFiles(writeFile func(filename string, content string) error, filesToCreate []internal.FileToCreate) (string, error) {
	for _, ftc := range filesToCreate {
		var content = ftc.Contents()
		var fullFileName = ftc.Path()
		var createHookFileError = writeFile(fullFileName, content)

		if createHookFileError != nil {
			return "", createHookFileError
		}
	}

	return hooksInstalledMessage, nil
}

var hooksInstalledMessage = "Hooks installed!"
