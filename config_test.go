package main

import "testing"

// Sum sums two numbers
func Sum(x int, y int) int {
	return x + y
}

func TestSum(t *testing.T) {
	var configuration = getConfiguration()

	if configuration.DoNotAutoAddHooksFromLocalHookDir != false {
		t.Errorf("Expected DoNotAutoAddHooksFromLocalHookDir to be false")
	}

	if configuration.LocalHookDir != "./hooks" {
		t.Errorf("LocalHookDir not as expected: %s", configuration.LocalHookDir)
	}

	var amountOfHooksFound = len(configuration.Hooks)
	if amountOfHooksFound != 1 {
		t.Errorf("Expected 1 Hook, found %d", amountOfHooksFound)
	}
}
