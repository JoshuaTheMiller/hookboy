package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

var hookConfigurationFileForStatementInstall = `
autoAddHooks: No
hooks:
  - hookName: theActualHookName        
    statement: echo This file was placed as a test!
    files: [] 
`

func TestThatHookStatementsGetInstalledProperly(t *testing.T) {
	var hookName = "post-update"
	var modifiedConfig = strings.Replace(hookConfigurationFileForStatementInstall, "theActualHookName", hookName, 1)

	// other tests test for configuration
	var configuration, configError = deserializeConfiguration([]byte(modifiedConfig))

	if configError != nil {
		t.Errorf("Test config is garbage: '%s'", configError)
		return
	}

	configuration.Install()

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

var hookConfigurationFileForLocalHookInstall = `
localHookDir: theActualLocalHookDir
autoAddHooks: ByFileName
hooks: []
`

func TestThatLocalHooksGetInstalledProperly(t *testing.T) {
	var hookName = "post-update"
	var localHookDir = "hooksForTesting"
	var modifiedConfig = strings.Replace(hookConfigurationFileForLocalHookInstall, "theActualLocalHookDir", localHookDir, 1)

	var fileContent = "echo Hello from testing!"
	os.Mkdir(localHookDir, 0755)
	var fileName = fmt.Sprintf("%s/%s", localHookDir, hookName)
	var mode = int(0777)
	var fileWriteError = ioutil.WriteFile(fileName, []byte(fileContent), os.FileMode(mode))

	if fileWriteError != nil {
		t.Errorf("Error writing test file: '%s'", fileWriteError)
		return
	}

	// other tests test for configuration
	var configuration, configError = deserializeConfiguration([]byte(modifiedConfig))

	if configError != nil {
		t.Errorf("Test config is garbage: '%s'", configError)
		return
	}

	configuration.Install()

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
