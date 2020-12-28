package aply

import (
	"io/ioutil"
	"os"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/prep"
)

func GetApplier() Applier {
	return &applierboy{}
}

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
	ab.instantiate()

	var prepboy = prep.Instantiate(configuration)

	var filesToCreate, _ = prepboy.PrepareHookfileInfo(configuration)

	return writeFiles(ab.WriteFile, filesToCreate)
}

func writeFiles(writeFile func(filename string, content string) error, filesToCreate []prep.FileToCreate) (string, error) {
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
