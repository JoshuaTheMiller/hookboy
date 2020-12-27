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

type simpleFile interface {
	Name() string
}

type applierboy struct {
	FilesToCreate map[string][]string
	ReadDir       func(dirname string) ([]simpleFile, error)
	WriteFile     func(filename string, content string) error
}

func readDir(dir string) ([]simpleFile, error) {
	var files, err = ioutil.ReadDir(dir)

	data := make([]simpleFile, len(files))

	for i := range files {
		data[i] = simpleFile(files[i])
	}

	return data, err
}

func writeFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), os.ModePerm)
}

func (ab *applierboy) instantiate() {
	ab.FilesToCreate = make(map[string][]string)
	ab.ReadDir = readDir
	ab.WriteFile = writeFile
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
			filePath, err := ab.generateStatementFile(hook.HookName, hook.Statement, configuration)

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

	return writeHookFiles(ab.WriteFile, ab.FilesToCreate, conf.ActualGitHooksDir)
}

func writeHookFiles(writeFile func(filename string, content string) error, filesToCreate map[string][]string, baseDir string) (string, error) {
	for fileName, linesForFile := range filesToCreate {
		var content = generateHookFile(linesForFile)
		var fullFileName = baseDir + "/" + fileName
		var createHookFileError = writeFile(fullFileName, content)

		if createHookFileError != nil {
			return "", createHookFileError
		}
	}

	return hooksInstalledMessage, nil
}

func (ab *applierboy) addHooksByFileName(localHooksDir string) error {
	files, err := ab.ReadDir(localHooksDir)

	if err != nil {
		return err
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

	return nil
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

var fileTemplateString = `#!/bin/sh
insertLinesHere

if insertConditionalHere;
then
exit 1
fi
exit 0`

func generateHookFile(linesToAdd []string) string {
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

func (ab applierboy) generateStatementFile(fileName string, statement string, conf conf.Configuration) (string, error) {
	var cacheDir = conf.GetCacheDirectory()
	os.MkdirAll(cacheDir, os.ModePerm)
	var filePath = cacheDir + "/" + fileName + "-statement"

	var err = ab.WriteFile(filePath, statement)

	if err != nil {
		return "", err
	}

	return filePath, nil
}

// HooksInstalledMessage
var hooksInstalledMessage = "Hooks installed!"
