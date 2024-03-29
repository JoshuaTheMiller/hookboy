package source

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
)

type localFolderReader struct {
	Path string
	Desc string
}

func (l localFolderReader) CanRead() bool {
	result := checkForFileSystemObject(l.Path)

	if result == doesNotExist || result == file {
		return false
	}

	return true
}

func (l localFolderReader) Read() (conf.Configuration, hookboy.Error) {
	var configurationToReturn conf.Configuration
	configurationToReturn.AutoAddHooks = conf.ByFileName
	configurationToReturn.LocalHookDir = l.Path

	return configurationToReturn, nil
}

func (l localFolderReader) Description() string {
	return l.Desc
}

func (l localFolderReader) Location() string {
	return l.Path
}
