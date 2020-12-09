package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

type extraArguments struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type hookFile struct {
	Path           string           `yaml:"path"`
	ExtraArguments []extraArguments `yaml:"extraArguments"`
}

type hooks struct {
	HookName string     `yaml:"hookName"`
	Files    []hookFile `yaml:"files"`
}

type configuration struct {
	LocalHookDir                      string  `yaml:"localHookDir"`
	DoNotAutoAddHooksFromLocalHookDir bool    `yaml:"doNotAutoAddHooksFromLocalHookDir"`
	Hooks                             []hooks `yaml:"hooks"`
}

func getConfiguration() *configuration {

	yamlFile, err := ioutil.ReadFile(".gitgrapple.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	c := &configuration{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c.setDefaults()
}

func (c *configuration) setDefaults() *configuration {
	if c.LocalHookDir == "" {
		c.LocalHookDir = "./hooks"
	}

	return c
}

var recognizedHooks = [...]string{
	"applypatch-msg",
	"commit-msg",
	"fsmonitor-watchman",
	"post-update",
	"pre-applypatch",
	"pre-commit",
	"pre-merge-commit",
	"pre-push",
	"pre-rebase",
	"pre-receive",
	"prepare-commit-msg",
	"update"}

var localGitHooksDir = ".git/hooks/"

func main() {
	var configuration = getConfiguration()

	var localHooksDir = configuration.LocalHookDir

	files, err := ioutil.ReadDir(localHooksDir)

	if err != nil {
		log.Fatal(err)
	}

	var filesToCreate = make(map[string][]string)

	for _, hook := range configuration.Hooks {
		lines := []string{}
		for _, fileToInclude := range hook.Files {
			lines = append(lines, generateLineFromFile(fileToInclude))
		}

		filesToCreate[hook.HookName] = lines
	}

	if !configuration.DoNotAutoAddHooksFromLocalHookDir {
		for _, f := range files {
			var potentialHookName = f.Name()

			if itemExists(recognizedHooks, potentialHookName) {
				var execLine = "exec \"./hooks/" + potentialHookName + "\"" + " \"$@\" "

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
		createBashExecFile(fileName, linesForFile)
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

func generateLineFromFile(fileToInclude hookFile) string {
	var sb strings.Builder

	sb.WriteString("exec ")
	sb.WriteString("\"" + fileToInclude.Path + "\" \"$@\" ")

	for _, arg := range fileToInclude.ExtraArguments {
		sb.WriteString(arg.Name + "=" + arg.Value + " ")
	}

	return sb.String()
}

func createBashExecFile(fileName string, linesToAdd []string) {
	file, err := os.Create(localGitHooksDir + "/" + fileName)

	if err != nil {
		log.Fatal(err)
		return
	}

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

	if err2 != nil {
		log.Fatal(err2)
	}
}
