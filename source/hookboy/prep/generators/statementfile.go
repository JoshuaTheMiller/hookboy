package generators

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

func prepareStatementFile(hookname string, statement string, c conf.Configuration) internal.FileToCreate {
	var cacheDir = c.GetCacheDirectory()
	var filePath = cacheDir + "/" + hookname + "-statement"

	return internal.FileToCreate{
		Path:     filePath,
		Contents: statement,
	}
}
