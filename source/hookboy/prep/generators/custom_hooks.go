package generators

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type customHookGenerator struct {
}

func (chg customHookGenerator) Generate(c conf.Configuration, readDir func(dirname string) ([]p.SimpleFile, error)) (ef []p.ExecutableFile, ftc []internal.FileToCreate, err error) {
	for _, hook := range c.Hooks {
		for _, fileToInclude := range hook.Files {
			var file = p.ExecutableFile{
				AssociatedHook: hook.HookName,
				Path:           fileToInclude.Path,
				ExtraArguments: fileToInclude.ExtraArguments,
			}

			ef = append(ef, file)
		}

		// This check, and the generation, should be moved to a sub custom hook generator
		if statementIsPresent(hook.Statement) {
			var statementFile = prepareStatementFile(hook.HookName, hook.Statement, c)

			ftc = append(ftc, statementFile)

			ef = append(ef, p.ExecutableFile{
				AssociatedHook: hook.HookName,
				Path:           statementFile.Path,
			})
		}
	}

	return ef, ftc, err
}

func (chg customHookGenerator) Name() string {
	return "Custom Hooks"
}

func statementIsPresent(s string) bool {
	return s != ""
}
