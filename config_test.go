package main

import (
	"fmt"
	"testing"
)

// Sum sums two numbers
func Sum(x int, y int) int {
	return x + y
}

func TestSum(t *testing.T) {
	var configuration, _ = getDefaultConfiguration()

	if configuration.AutoAddHooks != byFileName {
		t.Errorf("Expected AutoAddHooks to be byFileName")
	}

	if configuration.LocalHookDir != "./hooks" {
		t.Errorf("LocalHookDir not as expected: %s", configuration.LocalHookDir)
	}

	var amountOfHooksFound = len(configuration.Hooks)
	if amountOfHooksFound != 1 {
		t.Errorf("Expected 1 Hook, found %d", amountOfHooksFound)
	}
}

var testData1 = `
localHookDir: ./somethingElse
autoAddHooks: No
hooks: []
`

func TestAutoAddHooksSetToNoneProperly(t *testing.T) {
	var configuration, _ = deserializeConfiguration([]byte(testData1))

	if configuration.AutoAddHooks != no {
		t.Errorf("Expected AutoAddHooks to be No")
	}

	var expectedLocalHooksDir = "./somethingElse"
	if configuration.LocalHookDir != expectedLocalHooksDir {
		t.Errorf("LocalHookDir, expected %s, got %s", expectedLocalHooksDir, configuration.LocalHookDir)
	}
}

var testData2 = `
localHookDir: []
autoAddHooks: What
hooks: []
`

func TestInvalidConfigDataTreatedAsError(t *testing.T) {
	var _, deserializationError = deserializeConfiguration([]byte(testData2))

	var expectedErrorMessage = "failed to parse configuration file"
	var actualErrorMessage = deserializationError.Error()
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message to be '%s', got '%s'", expectedErrorMessage, actualErrorMessage)
	}
}

func TestInvalidConfigFileTreatedAsError(t *testing.T) {
	var fakeFileName = "somefilethatisnotreal"
	var _, fileReadError = getConfiguration(fakeFileName)

	var expectedErrorMessage = fmt.Sprintf("cannot read file '%s', please check that it is valid", fakeFileName)
	var actualErrorMessage = fileReadError.Error()
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message to be '%s', got '%s'", expectedErrorMessage, actualErrorMessage)
	}
}
