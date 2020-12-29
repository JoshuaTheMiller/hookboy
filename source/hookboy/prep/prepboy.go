package prep

import (
	"io/ioutil"

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

	ef, ftc, err := generateExecutableFilesAndFilesToCreate(c, p.ReadDir)

	if err != nil {
		return nil, err
	}

	var hookfilesToCreate = generateHookFileContents(ef, c)

	ftc = append(ftc, hookfilesToCreate...)

	return ftc, nil
}

func generateExecutableFilesAndFilesToCreate(c conf.Configuration, readDir func(dirname string) ([]simpleFile, error)) (ef []executableFile, ftc []FileToCreate, err error) {
	for _, hook := range c.Hooks {

		for _, fileToInclude := range hook.Files {
			var file = executableFile{
				AssociatedHook: hook.HookName,
				Path:           fileToInclude.Path,
				ExtraArguments: fileToInclude.ExtraArguments,
			}

			ef = append(ef, file)
		}

		if statementIsPresent(hook.Statement) {
			var statementFile = prepareStatementFile(hook.HookName, hook.Statement, c)

			ftc = append(ftc, statementFile)

			ef = append(ef, executableFile{
				AssociatedHook: hook.HookName,
				Path:           statementFile.Path(),
			})
		}
	}

	if c.AutoAddHooks == conf.ByFileName {
		var files, err = getHooksByFileName(c.LocalHookDir, readDir)

		if err != nil {
			return nil, nil, prepboyError{
				Description:   "Error prepping hooks by filename",
				InternalError: err,
			}
		}

		ef = append(ef, files...)
	}

	return ef, ftc, nil
}

func statementIsPresent(s string) bool {
	return s != ""
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

type prepboyError struct {
	Description   string
	InternalError error
}

func (pbe prepboyError) Error() string {
	return pbe.Description
}
