package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

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

var customHooksSourceDir = "./hooks/"

var localGitHooksDir = ".git/hooks/"

var pathToHookExecutionTemplate = "MyNameWillChange.sh"

func main() {
	files, err := ioutil.ReadDir(customHooksSourceDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if itemExists(recognizedHooks, f.Name()) {
			contents, err := ioutil.ReadFile(pathToHookExecutionTemplate)

			if err != nil {
				log.Fatal(err)
				return
			}

			templateContents := string(contents)
			var filledTemplate = strings.Replace(templateContents, "PathToRealHookFile", "\"./hooks/"+f.Name()+"\"", 1)

			file, err := os.Create(localGitHooksDir + "/" + f.Name())

			if err != nil {
				log.Fatal(err)
				return
			}

			_, err2 := file.WriteString(filledTemplate)

			if err2 != nil {
				log.Fatal(err2)
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
