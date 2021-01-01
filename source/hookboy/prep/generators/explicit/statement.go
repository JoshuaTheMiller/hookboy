package explicit

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type statement struct {
}

func (s statement) Name() string {
	return "statement"
}

func (s statement) Prepare(h conf.Hooks, c conf.Configuration) ([]p.ExecutableFile, []internal.FileToCreate, error) {
	if !statementIsPresent(h.Statement) {
		return []p.ExecutableFile{}, []internal.FileToCreate{}, nil
	}

	var statementFile = prepareStatementFile(h.HookName, h.Statement, c)

	return []p.ExecutableFile{p.ExecutableFile{
		AssociatedHook: h.HookName,
		Path:           statementFile.Path,
	}}, []internal.FileToCreate{statementFile}, nil
}

func prepareStatementFile(hookname string, statement string, c conf.Configuration) internal.FileToCreate {
	var cacheDir = c.GetCacheDirectory()
	var filePath = cacheDir + "/" + hookname + "-statement"

	return internal.FileToCreate{
		Path:     filePath,
		Contents: statement,
	}
}

func statementIsPresent(s string) bool {
	return s != ""
}
