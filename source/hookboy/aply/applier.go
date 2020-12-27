package aply

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/hookboy/source/hookboy/conf"
)

func GetApplier() Applier {
	return &applierboy{}
}

type Applier interface {
	Install(configuration conf.Configuration) (string, error)
}

type applierboy struct {
	FilesToCreate map[string][]string
}

func (ab *applierboy) instantiate() {
	ab.FilesToCreate = make(map[string][]string)
}

// Install installs the hooks with the given configuration
func (ab *applierboy) Install(configuration conf.Configuration) (string, error) {
	ab.instantiate()

	for _, hook := range configuration.Hooks {
		lines := []string{}
		for _, fileToInclude := range hook.Files {
			lines = append(lines, generateLineFromFile(fileToInclude))
		}

		if hook.Statement != "" {
			filePath, err := generateStatementFile(hook.HookName, hook.Statement, configuration)

			if err != nil {
				return "", err
			}

			var sb strings.Builder

			sb.WriteString("exec ")
			sb.WriteString("\"" + filePath + "\" \"$@\" ")

			lines = append(lines, sb.String())
		}

		ab.FilesToCreate[hook.HookName] = lines
	}

	if configuration.AutoAddHooks == conf.ByFileName {
		ab.addHooksByFileName(configuration.LocalHookDir)
	}

	for fileName, linesForFile := range ab.FilesToCreate {
		var createBashError = createBashExecFile(fileName, linesForFile)

		if createBashError != nil {
			return "", createBashError
		}
	}

	return hooksInstalledMessage, nil
}

func (ab *applierboy) addHooksByFileName(localHooksDir string) {
	files, err := ioutil.ReadDir(localHooksDir)

	if err != nil {
		// Should return actual error
		return
	}

	for _, f := range files {
		var potentialHookName = f.Name()

		if itemExists(conf.RecognizedHooks, potentialHookName) {
			execLine := "exec \"./localHooksDirToReplace/" + potentialHookName + "\"" + " \"$@\" "
			execLine = strings.Replace(execLine, "localHooksDirToReplace", localHooksDir, 1)

			currentLines, exists := ab.FilesToCreate[potentialHookName]

			if exists {
				currentLines = append(currentLines, execLine)
				ab.FilesToCreate[potentialHookName] = currentLines
			} else {
				var execLineAsArray = []string{execLine}
				ab.FilesToCreate[potentialHookName] = execLineAsArray
			}
		}
	}
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

func generateLineFromFile(fileToInclude conf.HookFile) string {
	var sb strings.Builder

	sb.WriteString("exec ")
	sb.WriteString("\"" + fileToInclude.Path + "\" \"$@\" ")

	for _, arg := range fileToInclude.ExtraArguments {
		sb.WriteString(arg.Name + "=" + arg.Value + " ")
	}

	return sb.String()
}

func createBashExecFile(fileName string, linesToAdd []string) error {
	file, err := os.Create(conf.ActualGitHooksDir + "/" + fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	var execFile = generateBashExecFile(linesToAdd)

	_, err2 := file.WriteString(execFile)

	return err2
}

var fileTemplateString = `#!/bin/sh
insertLinesHere

if insertConditionalHere;
then
exit 1
fi
exit 0`

func generateBashExecFile(linesToAdd []string) string {
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

	return withConditionalInserted
}

func generateStatementFile(fileName string, statement string, conf conf.Configuration) (string, error) {
	var cacheDir = conf.GetCacheDirectory()
	os.MkdirAll(cacheDir, os.ModePerm)
	var filePath = cacheDir + "/" + fileName + "-statement"
	file, err := os.Create(filePath)

	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err2 := file.WriteString(statement)

	return filePath, err2
}

// HooksInstalledMessage
var hooksInstalledMessage = "Hooks installed!"
