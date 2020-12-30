package aply

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/hookboy/source/hookboy/conf"
	_ "github.com/hookboy/source/hookboy/prep"
	_ "github.com/hookboy/source/hookboy/prep/generators"
)

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

	var applier = applierboy{}
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

	var applier = applierboy{}
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
		t.Errorf("Generated file incorrect. Expected \n'%s'\n, found \n'%s'\n", expectedContents, actualContents)
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
