package cli

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
	var testApplication FakeApplication
	var args = []string{"a", "hello"}
	var byteBuffer bytes.Buffer
	var err = RunApp(args, &byteBuffer, &testApplication)

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
	var installMessage = `Hooks installed!`

	var testApplication FakeApplication
	testApplication.InstallMessage = installMessage

	var args = []string{"a", "install"}
	var byteBuffer bytes.Buffer
	var err = RunApp(args, &byteBuffer, &testApplication)

	if err != nil {
		t.Errorf("Command failed to run: '%s'", err)
		return
	}

	var expectedOutput = installMessage
	var actualOutput = byteBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("Output incorrect! Expected '%s', received '%s'", expectedOutput, actualOutput)
	}
}

type FakeApplication struct {
	InstallMessage string
}

func (fakeApplication *FakeApplication) Install() (string, error) {
	return fakeApplication.InstallMessage, nil
}
