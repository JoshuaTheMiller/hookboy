package aply

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

type fileWriter func(filename string, data []byte, perm os.FileMode) error
type folderCreator func(path string, perm os.FileMode) error

type applierboy struct {
	instantiated bool
	writeFile    fileWriter
	createFolder folderCreator
}

// Install installs the hooks with the given configuration
func (ab applierboy) Install(configuration conf.Configuration, ftc []internal.FileToCreate) (string, error) {
	ab.instantiate()

	for _, ftc := range ftc {
		var content = ftc.Contents
		var fullFileName = ftc.Path

		var dir = path.Dir(fullFileName)

		err := ab.createFolder(dir, os.ModePerm)

		if err != nil {
			return "", internal.AplyError{
				Description:   "Problem creating new folder for file",
				InternalError: err,
			}
		}

		err = ab.writeFile(fullFileName, []byte(content), os.ModePerm)

		if err != nil {
			return "", internal.AplyError{
				Description:   "Problem creating file",
				InternalError: err,
			}
		}
	}

	return hooksInstalledMessage, nil
}

func (ab *applierboy) instantiate() {
	if !ab.instantiated {
		ab.writeFile = ioutil.WriteFile
		ab.createFolder = os.MkdirAll
		ab.instantiated = true
	}
}

var hooksInstalledMessage = "Hooks installed!"
