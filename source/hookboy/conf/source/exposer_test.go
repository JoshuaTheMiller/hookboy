package source

import (
	"testing"

	"github.com/hookboy/source/hookboy/internal"
)

func TestGetCurrentConfigurationSourceReturnsNoneWhenNoConfigurationIsPresent(t *testing.T) {
	var exposer = configurationExposer{}

	var _, err = exposer.LocateCurrentConfigurationSource()

	if err == nil {
		t.Error("Expected error to be returned, none was")
		return
	}

	if err != internal.NoConfigurationSourceFoundError {
		t.Errorf("Expected err to be 'NoConfigurationSourceFoundError, was: %s", err)
	}
}

func TestGetCurrentReturnsExpectedConfigurationSource(t *testing.T) {
	var fileToCreate = ".hookboy.yml"

	var contentsForTestGetCurrentExpectedConfigurationSource = `hooks:
- hookName: commit-msg
  statement: "echo 'Hi!`

	var fileSystemObject = fileSystemObjectOptions{
		Name:         fileToCreate,
		FileContents: contentsForTestGetCurrentExpectedConfigurationSource,
	}
	creatFileSystemObjectForTest(fileSystemObject)
	defer deleteFileSystemObjectForTest(fileSystemObject)

	var exposer = configurationExposer{}

	var configSource, err = exposer.LocateCurrentConfigurationSource()

	if err != nil {
		t.Errorf("Returned error when none was expected: %s", err)
		return
	}

	if configSource.Path != fileToCreate {
		t.Errorf("Expected config.Path to be %s, was %s", fileToCreate, configSource.Path)
		return
	}
}

// RetrieveCurrentConfiguration
var contentsForTestGetCurrentReturnsExpectedConfiguration = `---

hooks:
  - hookName: commit-msg
    statement: echo "Hi!"`

func TestGetCurrentReturnsExpectedConfiguration(t *testing.T) {
	var fileToCreate = ".hookboy.yml"

	var fileSystemObject = fileSystemObjectOptions{
		Name:         fileToCreate,
		FileContents: contentsForTestGetCurrentReturnsExpectedConfiguration,
	}
	creatFileSystemObjectForTest(fileSystemObject)
	defer deleteFileSystemObjectForTest(fileSystemObject)

	var exposer = configurationExposer{}

	var config, err = exposer.RetrieveCurrentConfiguration()

	if err != nil {
		t.Errorf("Returned error when none was expected: %s", err)
		return
	}

	var hooks = config.Hooks

	if len(hooks) != 1 {
		t.Errorf("Expected amount of hooks in configuration to be 1")
	}

	if hooks[0].HookName == "commit-msg" && hooks[0].Statement == "echo 'Hi!`" {
		t.Errorf("Hook from configuration file was not as expected:\n  HookName:%s\n  Statement:%s", hooks[0].HookName, hooks[0].Statement)
	}
}
