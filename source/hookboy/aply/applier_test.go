package aply

import (
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

type simpleFileForTest struct {
	name string
}

func (sf simpleFileForTest) Name() string {
	return sf.name
}
