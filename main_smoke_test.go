package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var testFolderName = ".test_bin"
var binaryName = "hookboy"

func getOsDependentCommandPath() string {
	var os = runtime.GOOS

	switch os {
	case "windows":
		return fmt.Sprintf("%s/%s.exe", testFolderName, binaryName)
	case "linux":
		return fmt.Sprintf("%s/%s", testFolderName, binaryName)
	default:
		var panicMessage = fmt.Sprintf("Operating system of '%s' not supported for acceptance tests!", os)
		panic(panicMessage)
	}
}

func TestMain(m *testing.M) {
	os.Mkdir(testFolderName, 0755)

	var commandToRun = exec.Command("go", "build", "-o", testFolderName)
	var runError = commandToRun.Run()

	if runError != nil {
		var panicMessage = fmt.Sprintf("Failed to prepare for tests!")
		panic(panicMessage)
	}

	var exitCode = m.Run()

	// TODO: Why does commenting out the following cause this test to successfully get past windows firewall?
	// var removeError = os.RemoveAll(testFolderName)

	// if removeError != nil {
	// 	var panicMessage = fmt.Sprintf("Failed to remove test folder: '%s'", testFolderName)
	// 	panic(panicMessage)
	// }

	os.Exit(exitCode)
}

func TestCliSaysHello(t *testing.T) {
	var flagToTest = "hello"
	var output, err = exec.Command(getOsDependentCommandPath(), flagToTest).Output()

	if err != nil {
		t.Errorf("Command failed to run: '%s'", err)
		return
	}

	var expectedOutput = `Hello! We hope you are enjoying Grapple!`
	var actualOutput = string(output)
	if actualOutput != expectedOutput {
		t.Errorf("Output incorrect! Expected '%s', received '%s'", expectedOutput, actualOutput)
	}
}

var configSourceResult = `| Configuration Source
|--> Source: .hookboy.yml
|--> Description: Hookboy will look for configuration in a yaml file named .hookboy`

func TestHookboyConfigSourceFunctions(t *testing.T) {
	var flagToTest = []string{"config", "source"}
	var output, err = exec.Command(getOsDependentCommandPath(), flagToTest...).Output()

	if err != nil {
		t.Errorf("Command failed to run: '%s'", err)
		return
	}

	var expectedOutput = configSourceResult
	var actualOutput = string(output)
	if actualOutput != expectedOutput {
		t.Errorf("Output incorrect! Expected\n'%s'\nreceived\n'%s'", expectedOutput, actualOutput)
	}
}

func TestHookboyInstallFunction(t *testing.T) {
	var flagToTest = []string{"install"}
	var output, err = exec.Command(getOsDependentCommandPath(), flagToTest...).Output()

	if err != nil {
		t.Errorf("Command failed to run: '%s'", err)
		return
	}

	var expectedOutput = "Hooks installed!"
	var actualOutput = string(output)
	if actualOutput != expectedOutput {
		t.Errorf("Output incorrect! Expected\n'%s'\nreceived\n'%s'", expectedOutput, actualOutput)
	}
}
