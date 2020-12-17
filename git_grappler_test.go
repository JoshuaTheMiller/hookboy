package main

import (
	"fmt"
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
