package hookboy

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// Install installs the hooks with the given configuration
func (configuration *Configuration) Install() (string, error) {
	var localHooksDir = configuration.LocalHookDir

	files, err := ioutil.ReadDir(localHooksDir)

	if err != nil {
		return "", errors.New("Unable to find localHooksDir")
	}

	var filesToCreate = make(map[string][]string)

	for _, hook := range configuration.Hooks {
		lines := []string{}
		for _, fileToInclude := range hook.Files {
			lines = append(lines, generateLineFromFile(fileToInclude))
		}

		if hook.Statement != "" {
			filePath, err := generateStatementFile(hook.HookName, hook.Statement)

			if err != nil {
				return "", err
			}

			var sb strings.Builder

			sb.WriteString("exec ")
			sb.WriteString("\"" + filePath + "\" \"$@\" ")

			lines = append(lines, sb.String())
		}

		filesToCreate[hook.HookName] = lines
	}

	if configuration.AutoAddHooks == ByFileName {
		for _, f := range files {
			var potentialHookName = f.Name()

			if itemExists(recognizedHooks, potentialHookName) {
				execLine := "exec \"./localHooksDirToReplace/" + potentialHookName + "\"" + " \"$@\" "
				execLine = strings.Replace(execLine, "localHooksDirToReplace", localHooksDir, 1)

				currentLines, exists := filesToCreate[potentialHookName]

				if exists {
					currentLines = append(currentLines, execLine)
					filesToCreate[potentialHookName] = currentLines
				} else {
					var execLineAsArray = []string{execLine}
					filesToCreate[potentialHookName] = execLineAsArray
				}
			}
		}
	}

	for fileName, linesForFile := range filesToCreate {
		var createBashError = createBashExecFile(fileName, linesForFile)

		if createBashError != nil {
			return "", createBashError
		}
	}

	return hooksInstalledMessage, nil
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func generateLineFromFile(fileToInclude HookFile) string {
	var sb strings.Builder

	sb.WriteString("exec ")
	sb.WriteString("\"" + fileToInclude.Path + "\" \"$@\" ")

	for _, arg := range fileToInclude.ExtraArguments {
		sb.WriteString(arg.Name + "=" + arg.Value + " ")
	}

	return sb.String()
}

func createBashExecFile(fileName string, linesToAdd []string) error {
	file, err := os.Create(actualGitHooksDir + "/" + fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	var fileTemplateString = `#!/bin/sh
insertLinesHere

if insertConditionalHere;
then
exit 1
fi
exit 0`

	var formattedLinesToAdd strings.Builder
	for index, line := range linesToAdd {
		var line = fmt.Sprintf("retVal%d=%s\nretVal%d=$?\n", index, line, index)

		formattedLinesToAdd.WriteString(line)
	}

	var formattedInnerConditional strings.Builder
	var amountOfLinesAdded = len(linesToAdd)
	for i := 0; i < amountOfLinesAdded; i++ {
		var formattedConditionalPart = fmt.Sprintf("[ $retVal%d -ne 0 ]", i)
		formattedInnerConditional.WriteString(formattedConditionalPart)

		if i < amountOfLinesAdded-1 {
			formattedInnerConditional.WriteString(" || ")
		}
	}

	var withTextInserted = strings.Replace(fileTemplateString, "insertLinesHere", formattedLinesToAdd.String(), 1)
	var withConditionalInserted = strings.Replace(withTextInserted, "insertConditionalHere", formattedInnerConditional.String(), 1)

	_, err2 := file.WriteString(withConditionalInserted)

	return err2
}

func generateStatementFile(fileName string, statement string) (string, error) {
	os.MkdirAll(grappleCacheDir, os.ModePerm)
	var filePath = grappleCacheDir + "/" + fileName + "-statement"
	file, err := os.Create(filePath)

	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err2 := file.WriteString(statement)

	return filePath, err2
}
