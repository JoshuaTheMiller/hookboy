package prep

import "github.com/hookboy/source/hookboy/conf"

func prepareStatementFile(hookname string, statement string, c conf.Configuration) FileToCreate {
	var cacheDir = c.GetCacheDirectory()
	var filePath = cacheDir + "/" + hookname + "-statement"

	return fileToCreate{
		path:     filePath,
		contents: statement,
	}
}
