package main

import "testing"

// Sum sums two numbers
func Sum(x int, y int) int {
	return x + y
}

func TestSum(t *testing.T) {
	var configuration = getDefaultConfiguration()

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
	var configuration = deserializeConfiguration([]byte(testData1))

	if configuration.AutoAddHooks != no {
		t.Errorf("Expected AutoAddHooks to be No")
	}

	var expectedLocalHooksDir = "./somethingElse"
	if configuration.LocalHookDir != expectedLocalHooksDir {
		t.Errorf("LocalHookDir, expected %s, got %s", expectedLocalHooksDir, configuration.LocalHookDir)
	}
}
