package aply

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

func Test_applier_InstallWritesFilesToFolderThatDidNotExistBefore(t *testing.T) {
	c := conf.Configuration{}
	folder := "somerandomfolder"
	filepath := fmt.Sprintf("%s/somePath.txt", folder)
	ftc := internal.FileToCreate{
		Contents: "somecontents",
		Path:     filepath,
	}

	var applier = applierboy{}

	_, err := applier.Install(c, []internal.FileToCreate{ftc})

	if err != nil {
		t.Error("Got error when none was expected")
		return
	}

	contentBytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		t.Error("Got error when none was expected. File doesn't appear to exist in expected location")
		return
	}

	expectedContents := ftc.Contents
	actualContents := string(contentBytes)
	if expectedContents != actualContents {
		t.Errorf("Generated file incorrect. Expected '%s', found '%s'", expectedContents, actualContents)
	}

	var fileRemoveError = os.RemoveAll(folder)

	if fileRemoveError != nil {
		t.Errorf("Test cleanup failed! Unable to remove file at '%s': '%s'", folder, fileRemoveError.Error())
	}
}

func Test_applier_InstallPropagatesFileError(t *testing.T) {
	c := conf.Configuration{}
	ftc := internal.FileToCreate{
		Contents: "DoesNotMatter",
		Path:     "DoesNotMatter",
	}

	errorToReturn := errors.New("File writing failed")
	errorReturningWrite := func(string, []byte, os.FileMode) error {
		return errorToReturn
	}
	nonErrorReturningFolderCreate := func(path string, perm os.FileMode) error {
		return nil
	}
	var applier = applierboy{
		writeFile:    errorReturningWrite,
		createFolder: nonErrorReturningFolderCreate,
		instantiated: true,
	}

	_, err := applier.Install(c, []internal.FileToCreate{ftc})

	aplyError, ok := err.(internal.AplyError)

	if !ok {
		t.Error("Expected error to be of type AplyError")
		return
	}

	expectedMessage := "Problem creating file"
	actualMessage := aplyError.Error()
	if expectedMessage != actualMessage {
		t.Error("AplyError does not have the correct error message")
		return
	}
}

func Test_applier_InstallPropagatesFolderError(t *testing.T) {
	c := conf.Configuration{}
	ftc := internal.FileToCreate{
		Contents: "DoesNotMatter",
		Path:     "DoesNotMatter",
	}

	errorToReturn := errors.New("DoesNotMatter")
	errorReturningFolderCreate := func(path string, perm os.FileMode) error {
		return errorToReturn
	}
	var applier = applierboy{
		createFolder: errorReturningFolderCreate,
		instantiated: true,
	}

	_, err := applier.Install(c, []internal.FileToCreate{ftc})

	aplyError, ok := err.(internal.AplyError)

	if !ok {
		t.Error("Expected error to be of type AplyError")
		return
	}

	expectedMessage := "Problem creating new folder for file"
	actualMessage := aplyError.Error()
	if expectedMessage != actualMessage {
		t.Error("AplyError does not have the correct error message")
		return
	}
}
