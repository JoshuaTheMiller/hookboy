package prep

import (
	"fmt"
	"strings"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

var fileTemplateString = `#!/bin/sh

if [ ! -z "${hookboy_skip}" ] && [ ${hookboy_skip} = "true" ]
then
LGray='\x1b[41m'
NC='\x1b[0m'
echo -e "${LGray}Environment Variable 'hookboy_skip' is set, skipping hook validation!${NC}"
exit 0
fi

insertLinesHere

if insertConditionalHere;
then
exit 1
fi
exit 0`

func generateHookFileContents(ef []p.ExecutableFile, c conf.Configuration) []internal.FileToCreate {
	var filesGroupByHook = groupByHook(ef)

	var files = []internal.FileToCreate{}

	for key, values := range filesGroupByHook {
		var hookname = key
		var executableFiles = values

		var ftc = internal.FileToCreate{
			Path:     getHookFilePath(hookname, c),
			Contents: generateHookFileContent(executableFiles),
		}

		files = append(files, ftc)
	}

	return files
}

func getHookFilePath(hookname string, c conf.Configuration) string {
	return conf.ActualGitHooksDir + "/" + hookname
}

func generateHookFileContent(ef []p.ExecutableFile) string {
	var linesToAdd = generateExecutableLines(ef)

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

func generateExecutableLines(ef []p.ExecutableFile) []string {
	var execLines = []string{}
	for _, f := range ef {
		var line = generateLineFromFile(f)

		execLines = append(execLines, line)
	}

	return execLines
}

func generateLineFromFile(f p.ExecutableFile) string {
	var sb strings.Builder

	sb.WriteString("exec ")
	sb.WriteString("\"" + f.Path + "\" \"$@\" ")

	for _, arg := range f.ExtraArguments {
		sb.WriteString(arg.Name + "=" + arg.Value + " ")
	}

	return sb.String()
}

func groupByHook(ef []p.ExecutableFile) map[string][]p.ExecutableFile {
	var m = make(map[string][]p.ExecutableFile)

	for _, item := range ef {
		var key = item.AssociatedHook
		var value = item

		existingItems, exists := m[key]

		if exists {
			existingItems = append(existingItems, value)
			m[key] = existingItems
		} else {
			m[key] = []p.ExecutableFile{value}
		}
	}

	return m
}
