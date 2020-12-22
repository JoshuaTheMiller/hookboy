package aply

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

func TestGenerateExpectedLineFromFile(t *testing.T) {

	var extraArgNameValue = "extraArgName"
	var extraArgValueValue = "extraArgValue"
	var ea = conf.ExtraArguments{
		Name:  extraArgNameValue,
		Value: extraArgValueValue,
	}

	var pathValue = "somePath"
	var hookFile = conf.HookFile{
		Path:           pathValue,
		ExtraArguments: []conf.ExtraArguments{ea},
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

func TestGenerateStatementFileIsAsExpected(t *testing.T) {
	var statementName = "SomeStatement"
	var someStatement = "SomeStatement"

	generateStatementFile(statementName, someStatement)

	// After testing this, it is obviously not clear that the func above
	// will generate a file at the path below...
	var filePath = fmt.Sprintf(".grapple-cache/%s-statement", statementName)
	var contentBytes, error = ioutil.ReadFile(filePath)

	if error != nil {
		t.Errorf("File generation appears broken: %s", error.Error())
		return
	}

	var actualContents = string(contentBytes)
	var expectedContents = someStatement

	if expectedContents != actualContents {
		t.Errorf("Generated file incorrect. Expected '%s', found '%s'", expectedContents, actualContents)
	}

	var fileRemoveError = os.Remove(filePath)

	if fileRemoveError != nil {
		t.Errorf("Test cleanup failed! Unable to remove file at '%s': '%s'", filePath, fileRemoveError.Error())
	}
}

func TestItemExists(t *testing.T) {
	var someItem = "what"
	var someArray = [...]string{someItem, "what2"}

	// itemExists is not very clear as far as what param goes where
	// #CopyPasteFails
	var itemExists = itemExists(someArray, someItem)

	if !itemExists {
		t.Error("Item existed, but not found")
	}
}

func TestThatHookStatementsGetInstalledProperly(t *testing.T) {
	var hookName = "post-update"

	// other tests test for configuration
	var configuration = conf.Configuration{
		AutoAddHooks: conf.No,
		Hooks: []conf.Hooks{
			conf.Hooks{
				HookName:  "theActualHookName",
				Statement: "echo This file was placed as a test!",
			},
		},
	}

	Install(configuration)

	var filePath = fmt.Sprintf(".git/hooks/%s", hookName)
	var contentBytes, error = ioutil.ReadFile(filePath)

	if error != nil {
		t.Errorf("File generation appears broken: %s", error.Error())
		return
	}

	var actualContents = string(contentBytes)
	var expectedContents = `#!/bin/sh
retVal0=exec ".grapple-cache/post-update-statement" "$@" 
retVal0=$?


if [ $retVal0 -ne 0 ];
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

func TestThatLocalHooksGetInstalledProperly(t *testing.T) {
	var hookName = "post-update"
	var localHookDir = "hooksForTesting"

	var fileContent = "echo Hello from testing!"
	os.Mkdir(localHookDir, 0777)
	var fileName = fmt.Sprintf("%s/%s", localHookDir, hookName)
	var mode = int(0777)
	var fileWriteError = ioutil.WriteFile(fileName, []byte(fileContent), os.FileMode(mode))

	if fileWriteError != nil {
		t.Errorf("Error writing test file: '%s'", fileWriteError)
		return
	}

	// other tests test for configuration
	var configuration = conf.Configuration{
		LocalHookDir: "theActualLocalHookDir",
		AutoAddHooks: conf.ByFileName,
	}

	Install(configuration)

	var filePath = fmt.Sprintf(".git/hooks/%s", hookName)
	var contentBytes, error = ioutil.ReadFile(filePath)

	if error != nil {
		t.Errorf("File generation appears broken: %s", error.Error())
		return
	}

	var actualContents = string(contentBytes)
	expectedContents := `#!/bin/sh
retVal0=exec "./theActualLocalHookDir/post-update" "$@" 
retVal0=$?


if [ $retVal0 -ne 0 ];
then
exit 1
fi
exit 0`
	expectedContents = strings.Replace(expectedContents, "theActualLocalHookDir", localHookDir, 1)

	if expectedContents != actualContents {
		t.Errorf("Generated file incorrect. Expected '%s', found '%s'", expectedContents, actualContents)
	}

	var testFolderCleanupError = os.RemoveAll(localHookDir)
	if testFolderCleanupError != nil {
		t.Errorf("Test cleanup failed! Unable to remove folder at '%s': '%s'", localHookDir, testFolderCleanupError.Error())
	}

	var pathToHook = fmt.Sprintf(".git/hooks/%s", hookName)
	var fileRemoveError = os.Remove(pathToHook)

	if fileRemoveError != nil {
		t.Errorf("Test cleanup failed! Unable to remove file at '%s': '%s'", pathToHook, fileRemoveError.Error())
	}
}
