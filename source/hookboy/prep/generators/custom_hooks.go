package generators

import (
	"net/url"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type customHookGenerator struct {
}

func (chg customHookGenerator) Generate(c conf.Configuration, readDir func(dirname string) ([]p.SimpleFile, error)) (ef []p.ExecutableFile, ftc []internal.FileToCreate, err error) {
	for _, hook := range c.Hooks {
		// TODO: seems like a pattern has emerged here... Abstract?
		var localExecutableFiles, localFilesToCreate, localPrepareError = prepareLocalHooks(hook, c)

		if localPrepareError != nil {
			return nil, nil, localPrepareError
		}

		ef = append(ef, localExecutableFiles...)
		ftc = append(ftc, localFilesToCreate...)

		var nonLocalExecutableFiles, nonLocalFilesToCreate, nonLocalPrepareError = prepareNonLocalHooks(hook, c)

		if nonLocalPrepareError != nil {
			return nil, nil, nonLocalPrepareError
		}

		ef = append(ef, nonLocalExecutableFiles...)
		ftc = append(ftc, nonLocalFilesToCreate...)

		var statementExecutableFiles, statementPreparedFilesToCreate, statementError = prepareStatement(hook, c)

		if statementError != nil {
			return nil, nil, statementError
		}

		ef = append(ef, statementExecutableFiles...)
		ftc = append(ftc, statementPreparedFilesToCreate...)
	}

	return ef, ftc, err
}

func prepareLocalHooks(hook conf.Hooks, c conf.Configuration) (ef []p.ExecutableFile, ftc []internal.FileToCreate, err error) {
	for _, fileToInclude := range hook.Files {
		var path = fileToInclude.Path

		if pathIsNonLocalPath(path) {
			// This func only handles local files, skip if non-local
			continue
		}

		var file = p.ExecutableFile{
			AssociatedHook: hook.HookName,
			Path:           fileToInclude.Path,
			ExtraArguments: fileToInclude.ExtraArguments,
		}

		ef = append(ef, file)
	}

	return ef, ftc, nil
}

var nonLocalFileError = internal.PrepError{
	Description: "Non-local files are not yet supported!",
}

func prepareNonLocalHooks(hook conf.Hooks, c conf.Configuration) (ef []p.ExecutableFile, ftc []internal.FileToCreate, err error) {
	for _, fileToInclude := range hook.Files {
		var path = fileToInclude.Path

		if !pathIsNonLocalPath(path) {
			// This func only handles non-local files, skip if not non-local
			continue
		}

		return nil, nil, nonLocalFileError
	}

	return ef, ftc, nil
}

func prepareStatement(hook conf.Hooks, c conf.Configuration) ([]p.ExecutableFile, []internal.FileToCreate, error) {
	// This check, and the generation, should be moved to a sub custom hook generator
	if !statementIsPresent(hook.Statement) {
		return []p.ExecutableFile{}, []internal.FileToCreate{}, nil
	}

	var statementFile = prepareStatementFile(hook.HookName, hook.Statement, c)

	ftc := statementFile

	ef := p.ExecutableFile{
		AssociatedHook: hook.HookName,
		Path:           statementFile.Path,
	}

	return []p.ExecutableFile{ef}, []internal.FileToCreate{ftc}, nil
}

func pathIsNonLocalPath(path string) bool {
	u, err := url.Parse(path)

	// TODO: this should most likely be expanded. There are edge cases this
	// will not catch. Decent starting point though.
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (chg customHookGenerator) Name() string {
	return "Custom Hooks"
}

func statementIsPresent(s string) bool {
	return s != ""
}
