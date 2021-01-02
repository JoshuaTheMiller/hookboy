package cli

import (
	"bytes"
	"testing"

	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/app"
	"github.com/hookboy/source/hookboy/conf"
)

func TestCliIsMappedToApplicationCorrectly(t *testing.T) {
	var commandMappingsToTest = map[string]commandMapTest{
		"hello": commandMapTest{
			ArgsToPass:     []string{"hello"},
			ExpectedOutput: "Hello! We hope you are enjoying Grapple!",
			Builder: builder{
				TestApplication: testApplication{},
			},
		},
		"install": commandMapTest{
			ArgsToPass:     []string{"install"},
			ExpectedOutput: "Installed!",
			Builder: builder{
				TestApplication: testApplication{
					InstallMessage: "Installed!",
				},
			},
		},
		"install-installError": commandMapTest{
			ArgsToPass:     []string{"install"},
			ExpectedOutput: "Installing failed",
			Builder: builder{
				TestApplication: testApplication{
					Error: hookboy.NewError("Installing failed"),
				},
			},
			IsErrorTest: true,
		},
		"install-buildError": commandMapTest{
			ArgsToPass:     []string{"install"},
			ExpectedOutput: "Building failed",
			Builder: builder{
				Error: hookboy.NewError("Building failed"),
			},
			IsErrorTest: true,
		},
		"config-source": commandMapTest{
			ArgsToPass:     []string{"config", "source"},
			ExpectedOutput: "SomeSourceLocation",
			Builder: builder{
				TestApplication: testApplication{
					ConfigLocation: "SomeSourceLocation",
				},
			},
			IsErrorTest: true,
		},
	}

	for _, test := range commandMappingsToTest {
		var expectedOutput = test.ExpectedOutput
		var args = formatArgsForCli(test.ArgsToPass...)

		var byteBuffer bytes.Buffer
		var err = RunApp(args, &byteBuffer, &test.Builder)

		var errorExists = err != nil

		if test.IsErrorTest && errorExists {
			if expectedOutput != err.Error() {
				t.Errorf("Command did not return expected error: %s", err)
			}
			continue
		}

		if errorExists {
			t.Errorf("Command failed to run: '%s'", err)
			continue
		}

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
	ArgsToPass     []string
	ExpectedOutput string
	Builder        builder
	IsErrorTest    bool
}

type builder struct {
	TestApplication testApplication
	Error           hookboy.Error
}

func (b *builder) Construct() (app.Application, hookboy.Error) {
	if b.Error != nil {
		return nil, b.Error
	}

	return &b.TestApplication, nil
}

type testApplication struct {
	InstallMessage string
	Error          hookboy.Error
	ConfigLocation string
	Config         conf.Configuration
}

func (ta *testApplication) Install() (string, hookboy.Error) {
	if ta.Error != nil {
		return "", ta.Error
	}

	return ta.InstallMessage, nil
}

func (ta *testApplication) CurrentConfiguration() (conf.Configuration, hookboy.Error) {
	return ta.Config, ta.Error
}

func (ta *testApplication) ConfigurationLocation() (string, hookboy.Error) {
	return ta.ConfigLocation, ta.Error
}
