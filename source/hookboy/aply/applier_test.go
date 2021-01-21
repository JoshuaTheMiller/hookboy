package aply

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	_, error := applier.Install(c, []internal.FileToCreate{ftc})

	if error != nil {
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
	writeCallCount := 0
	errorReturningWrite := func(string, []byte, os.FileMode) error {
		writeCallCount++

		if writeCallCount > 1 {
			return errorToReturn
		}

		return nil
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

	expectedMessage := "Problem creating file"
	actualMessage := err.Error()
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
	// How could I have written this such that the new code I wrote in this commit would
	// not break the previous tests and functionality? Decorator pattern could work
	nonErrorReturningFileWrite := func(string, []byte, os.FileMode) error {
		return nil
	}
	folderCreateCalledAmount := 0
	errorReturningFolderCreate := func(path string, perm os.FileMode) error {
		folderCreateCalledAmount++

		if folderCreateCalledAmount > 1 {
			return errorToReturn
		}

		return nil
	}
	var applier = applierboy{
		writeFile:    nonErrorReturningFileWrite,
		createFolder: errorReturningFolderCreate,
		instantiated: true,
	}

	_, err := applier.Install(c, []internal.FileToCreate{ftc})

	expectedMessage := "Problem creating new folder for file"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Error("AplyError does not have the correct error message")
		return
	}
}

func Test_writeInitialConfFile_GeneratesProperPathAndDefaultFile(t *testing.T) {
	expectedPath := ".git/hooks"
	expectedFileName := ".git/hooks/.hookboy-conf"
	expectedStartOfContents := `{
		"README": "These hooks have been wrangled by Hookboy!`

	actualFileName := ""
	actualContents := ""
	var fileWriter = func(filename string, data []byte, perm os.FileMode) error {
		actualFileName = filename
		actualContents = string(data)
		return nil
	}

	actualFolderPath := ""
	var folderCreator = func(path string, perm os.FileMode) error {
		actualFolderPath = path
		return nil
	}

	err := writeInitialConfFile(folderCreator, fileWriter)

	if err != nil {
		t.Error("Expected err to be nil")
		return
	}

	// Saying that this is "good enough"
	// Could deserialize back to the actual object to see if everything was written as expected I suppose
	if strings.HasPrefix(actualContents, expectedStartOfContents) {
		t.Error("Contents are not as expected")
	}

	if actualFolderPath != expectedPath {
		t.Error("Contents are not as expected")
	}

	if actualFileName != expectedFileName {
		t.Error("Contents are not as expected")
	}
}

type FileWriter func(filename string, data []byte, perm os.FileMode) error
type FolderCreator func(path string, perm os.FileMode) error
