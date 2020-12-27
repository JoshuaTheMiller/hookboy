package aply

import (
	"errors"
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

func TestGenerateStatementFileIsAsExpected(t *testing.T) {
	var statementName = "SomeStatement"
	var someStatement = "SomeStatement"

	var actualFileName = ""
	var actualContents = ""

	var conf = conf.Configuration{}
	var ab = applierboy{
		WriteFile: func(fileName string, content string) error {
			actualFileName = fileName
			actualContents = content

			return nil
		},
	}

	ab.generateStatementFile(statementName, someStatement, conf)

	var expectedContents = someStatement
	var expectedFileName = conf.GetCacheDirectory() + "/" + statementName + "-statement"

	if expectedContents != actualContents {
		t.Errorf("Generated file incorrect. Expected '%s', found '%s'", expectedContents, actualContents)
	}

	if expectedFileName != actualFileName {
		t.Errorf("Generated filename incorrect. Expected '%s', found '%s'", expectedFileName, actualFileName)
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
				HookName:  hookName,
				Statement: "echo This file was placed as a test!",
			},
		},
	}

	var applier = GetApplier()
	applier.Install(configuration)

	var filePath = fmt.Sprintf(".git/hooks/%s", hookName)
	var contentBytes, error = ioutil.ReadFile(filePath)

	if error != nil {
		t.Errorf("File generation appears broken: %s", error.Error())
		return
	}

	var actualContents = string(contentBytes)
	var expectedContents = `#!/bin/sh
retVal0=exec "/.hookboy-cache/post-update-statement" "$@" 
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
		LocalHookDir: localHookDir,
		AutoAddHooks: conf.ByFileName,
	}

	var applier = GetApplier()
	applier.Install(configuration)

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

type simpleFileForTest struct {
	name string
}

func (sf simpleFileForTest) Name() string {
	return sf.name
}

func TestAddHooksByFileNameAddsToExistingHooks(t *testing.T) {
	// must be a recongized hook for now
	var hookToAddTo = "commit-msg"

	var originalItems = map[string][]string{
		hookToAddTo: []string{"existing statement line for hook"},
	}

	var fileToAdd = hookToAddTo

	var applier = applierboy{
		FilesToCreate: originalItems,
		ReadDir: func(dirname string) ([]simpleFile, error) {
			return []simpleFile{
				simpleFileForTest{
					name: fileToAdd,
				},
			}, nil
		},
	}

	var err = applier.addHooksByFileName("doesnotmatter")

	if err != nil {
		t.Errorf("Found error when expected none: %s", err)
		return
	}

	if len(applier.FilesToCreate) != 1 {
		t.Errorf("Expected amount of hooks in list to be 1")
	}

	var files = applier.FilesToCreate[hookToAddTo]

	if len(files) != 2 {
		t.Errorf("Expected amount of hooks in list to be 2")
	}
}

func TestAddHooksReturnsEarlyOnError(t *testing.T) {
	var someError = errors.New("asdfasdf")

	var ab = applierboy{
		ReadDir: func(dirname string) ([]simpleFile, error) {
			return nil, someError
		},
	}

	// This doesn't really test that the method returns early, it just tests that the appropriate error is
	// returned.
	var actualError = ab.addHooksByFileName("doesNotMatterForThisTest")
	var expectedError = someError

	if expectedError != actualError {
		t.Errorf("Actual error not as expected. Expected '%s', received '%s'", expectedError, actualError)
	}
}
