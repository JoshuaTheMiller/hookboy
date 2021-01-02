package aply

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/hookboy/source/hookboy"
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
func (ab applierboy) Install(configuration conf.Configuration, ftc []internal.FileToCreate) (string, hookboy.Error) {
	ab.instantiate()

	for _, ftc := range ftc {
		var content = ftc.Contents
		var fullFileName = ftc.Path

		var dir = path.Dir(fullFileName)

		err := ab.createFolder(dir, os.ModePerm)

		if err != nil {
			return "", hookboy.WrapError(err, "Problem creating new folder for file")
		}

		err = ab.writeFile(fullFileName, []byte(content), os.ModePerm)

		if err != nil {
			return "", hookboy.WrapError(err, "Problem creating file")
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
