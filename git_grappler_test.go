package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateExpectedLineFromFile(t *testing.T) {

	var extraArgNameValue = "extraArgName"
	var extraArgValueValue = "extraArgValue"
	var ea = extraArguments{
		Name:  extraArgNameValue,
		Value: extraArgValueValue,
	}

	var pathValue = "somePath"
	var hookFile = hookFile{
		Path:           pathValue,
		ExtraArguments: []extraArguments{ea},
	}

	var expectedLine = fmt.Sprintf("exec \"%s\" \"$@\" %s=%s ", pathValue, extraArgNameValue, extraArgValueValue)
	var actualLine = generateLineFromFile(hookFile)

	if actualLine != expectedLine {
		t.Errorf("Generated Line Incorrect. Expected '%s', received '%s'", expectedLine, actualLine)
	}
}

// More of an Acceptance Test, as it is fairly high level and depends on the file system
func TestGeneratedFileIsAsExpected(t *testing.T) {
	var fileName = "samplefortest"
	var linesToAdd = []string{"line1", "line2"}
	createBashExecFile(fileName, linesToAdd)

	var filePath = fmt.Sprintf(".git/hooks/%s", fileName)
	var contentBytes, error = ioutil.ReadFile(filePath)

	if error != nil {
		t.Errorf("File generation appears broken: %s", error.Error())
	}

	var actualContents = string(contentBytes)
	var expectedContents = `#!/bin/sh
retVal0=line1
retVal0=$?
retVal1=line2
retVal1=$?


if [ $retVal0 -ne 0 ] || [ $retVal1 -ne 0 ];
then
exit 1
fi
exit 0`

	if expectedContents != actualContents {
		t.Errorf("Generated file incorrect. Expected '%s', found '%s'", expectedContents, actualContents)
	}

	var fileRemoveError = os.Remove(filePath)

	if fileRemoveError != nil {
		t.Errorf("Test cleanup failed! Unable to remove file at '%s': '%s'", filePath, fileRemoveError.Error())
	}
}
