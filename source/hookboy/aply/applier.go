package aply

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	"github.com/hookboy/source/internal/boundary"
)

type applierboy struct {
	instantiated bool
	writeFile    boundary.FileWriter
	createFolder boundary.FolderCreator
}

// Install installs the hooks with the given configuration
func (ab applierboy) Install(configuration conf.Configuration, ftc []internal.FileToCreate) (string, hookboy.Error) {
	ab.instantiate()

	err := prepareHookFolder(ab.createFolder, ab.writeFile)

	if err != nil {
		return "", hookboy.WrapError(err, "Error preparing git hooks directory")
	}

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

// TODO: too many responsibilities in this method
func prepareHookFolder(fc boundary.FolderCreator, fw boundary.FileWriter) error {
	// TODO: check for existance of ".hookboy-conf"
	// 	if found, update with last run information
	// 	if not found, and other files are present, zip other files
	//  Then, delete all files and folders besides ".hookboy-conf" and "snapshot-b4-hookboy" zip
	// 	Create ".hookboy-conf" if necessary
	// 	Then, run install as normal
	err := writeInitialConfFile(fc, fw)

	return err
}

func writeInitialConfFile(fc boundary.FolderCreator, fw boundary.FileWriter) error {
	var conf = hookboyConf{}
	conf.Default()

	niceJSON, err := boundary.SerializeToNiceJSON(conf)

	if err != nil {
		return err
	}

	var gitHooksDir = ".git/hooks"
	err = fc(gitHooksDir, os.ModePerm)

	if err != nil {
		return err
	}

	var fileName = fmt.Sprintf("%s/%s", gitHooksDir, filename)
	err = fw(fileName, []byte(niceJSON), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

var hooksInstalledMessage = "Hooks installed!"
