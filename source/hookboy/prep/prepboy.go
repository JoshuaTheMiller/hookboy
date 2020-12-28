package prep

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hookboy/source/hookboy/conf"
)

// Instantiate returns a Prepper for external use
func Instantiate(c conf.Configuration) Prepper {
	var p = prepboy{}
	p.instantiate(c)
	return p
}

func (p prepboy) PrepareHookfileInfo(c conf.Configuration) (ftc []FileToCreate, e error) {
	if p.Instantiated != true {
		p.instantiate(c)
	}

	var ef = []executableFile{}

	for _, hook := range c.Hooks {

		for _, fileToInclude := range hook.Files {
			var file = executableFile{
				AssociatedHook: hook.HookName,
				Path:           fileToInclude.Path,
				ExtraArguments: fileToInclude.ExtraArguments,
			}

			ef = append(ef, file)
		}

		if hook.Statement != "" {
			var statementFile = prepareStatementFile(hook.HookName, hook.Statement, c)

			ftc = append(ftc, statementFile)

			ef = append(ef, executableFile{
				AssociatedHook: hook.HookName,
				Path:           statementFile.Path(),
			})
		}
	}

	if c.AutoAddHooks == conf.ByFileName {
		var files, err = getHooksByFileName(c.LocalHookDir, p.ReadDir)

		if err != nil {
			return nil, prepboyError{
				Description:   "Error prepping hooks by filename",
				InternalError: err,
			}
		}

		ef = append(ef, files...)
	}

	var hookfilesToCreate = generateHookFileContents(ef, c)

	ftc = append(ftc, hookfilesToCreate...)

	return ftc, nil
}

// Prepper is used to prepare files to usage by Git Hook
type Prepper interface {
	PrepareHookfileInfo(c conf.Configuration) (ftc []FileToCreate, e error)
}

// TODO: Does this interface already exist?
type simpleFile interface {
	Name() string
}

type prepboy struct {
	Instantiated bool
	C            conf.Configuration
	ReadDir      func(dirname string) ([]simpleFile, error)
}

func (p *prepboy) instantiate(c conf.Configuration) {
	p.C = c
	p.ReadDir = readDir
	p.Instantiated = true
}

func readDir(dir string) ([]simpleFile, error) {
	var files, err = ioutil.ReadDir(dir)

	data := make([]simpleFile, len(files))

	for i := range files {
		data[i] = simpleFile(files[i])
	}

	return data, err
}

func generateHookFileContents(ef []executableFile, c conf.Configuration) []FileToCreate {
	var filesGroupByHook = groupByHook(ef)

	var files = []FileToCreate{}

	for key, values := range filesGroupByHook {
		var hookname = key
		var executableFiles = values

		var ftc = fileToCreate{
			path:     getHookFilePath(hookname, c),
			contents: generateHookFileContent(executableFiles),
		}

		files = append(files, ftc)
	}

	return files
}

func getHookFilePath(hookname string, c conf.Configuration) string {
	return conf.ActualGitHooksDir + "/" + hookname
}

func generateHookFileContent(ef []executableFile) string {
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

func generateExecutableLines(ef []executableFile) []string {
	var execLines = []string{}
	for _, f := range ef {
		var line = generateLineFromFile(f)

		execLines = append(execLines, line)
	}

	return execLines
}

func generateLineFromFile(f executableFile) string {
	var sb strings.Builder

	sb.WriteString("exec ")
	sb.WriteString("\"" + f.Path + "\" \"$@\" ")

	for _, arg := range f.ExtraArguments {
		sb.WriteString(arg.Name + "=" + arg.Value + " ")
	}

	return sb.String()
}

type executableFile struct {
	AssociatedHook string
	Path           string
	ExtraArguments []conf.ExtraArguments
}

type prepboyError struct {
	Description   string
	InternalError error
}

func (pbe prepboyError) Error() string {
	return pbe.Description
}

func groupByHook(ef []executableFile) map[string][]executableFile {
	var m = make(map[string][]executableFile)

	for _, item := range ef {
		var key = item.AssociatedHook
		var value = item

		existingItems, exists := m[key]

		if exists {
			existingItems = append(existingItems, value)
			m[key] = existingItems
		} else {
			m[key] = []executableFile{value}
		}
	}

	return m
}
