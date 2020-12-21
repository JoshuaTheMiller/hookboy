package runner

import (
	"bytes"
	"testing"
)

func formatArgsForCli(args ...string) []string {
	// The first value in the arg array is always the binary itself.
	// Without this, the call to app.Run will fail
	cliArguments := []string{"appName"}
	cliArguments = append(cliArguments, args...)
	return cliArguments
}

func TestRunAppSaysHello(t *testing.T) {
	// var flagToTest = "hello"
	var args = []string{"a", "hello"}
	var byteBuffer bytes.Buffer
	var err = RunApp(args, &byteBuffer)

	if err != nil {
		t.Errorf("Command failed to run: '%s'", err)
		return
	}

	var expectedOutput = `Hello! We hope you are enjoying Grapple!`
	var actualOutput = byteBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("Output incorrect! Expected '%s', received '%s'", expectedOutput, actualOutput)
	}
}

func TestRunAppInstallsSuccessfully(t *testing.T) {
	// var flagToTest = "hello"
	var args = []string{"a", "install"}
	var byteBuffer bytes.Buffer
	var err = RunApp(args, &byteBuffer)

	if err != nil {
		t.Errorf("Command failed to run: '%s'", err)
		return
	}

	var expectedOutput = `Hooks installed!`
	var actualOutput = byteBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("Output incorrect! Expected '%s', received '%s'", expectedOutput, actualOutput)
	}
}
