package cli

import (
	"bytes"
	"testing"
)

func TestCliIsMappedToApplicationCorrectly(t *testing.T) {
	var commandMappingsToTest = map[string]commandMapTest{
		"hello": commandMapTest{
			ArgsToPass:      []string{"hello"},
			ExpectedOutput:  "Hello! We hope you are enjoying Grapple!",
			FakeApplication: FakeApplication{},
		},
		"install": commandMapTest{
			ArgsToPass:     []string{"install"},
			ExpectedOutput: "Installed!",
			FakeApplication: FakeApplication{
				InstallMessage: "Installed!",
			},
		},
	}

	for _, test := range commandMappingsToTest {
		var testApplication = test.FakeApplication
		var args = formatArgsForCli(test.ArgsToPass...)

		var byteBuffer bytes.Buffer
		var err = RunApp(args, &byteBuffer, &testApplication)

		if err != nil {
			t.Errorf("Command failed to run: '%s'", err)
			return
		}

		var expectedOutput = test.ExpectedOutput
		var actualOutput = byteBuffer.String()
		if actualOutput != expectedOutput {
			t.Errorf("Output incorrect! Expected '%s', received '%s'", expectedOutput, actualOutput)
		}
	}
}

func formatArgsForCli(args ...string) []string {
	// The first value in the arg array is always the binary itself.
	// Without this, the call to app.Run will fail
	cliArguments := []string{"appName"}
	cliArguments = append(cliArguments, args...)
	return cliArguments
}

type commandMapTest struct {
	ArgsToPass      []string
	ExpectedOutput  string
	FakeApplication FakeApplication
}

type FakeApplication struct {
	InstallMessage string
}

func (fakeApplication *FakeApplication) Install() (string, error) {
	return fakeApplication.InstallMessage, nil
}
