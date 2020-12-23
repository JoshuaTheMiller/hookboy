package source

import "github.com/hookboy/source/hookboy/conf"

type localFileReader struct {
	Path string
	Desc string
}

func (l localFileReader) CanRead() bool {
	result := checkForFileSystemObject(l.Path)

	if result == doesNotExist || result == folder {
		return false
	}

	return true
}

func (l localFileReader) Read() (conf.Configuration, error) {
	var config conf.Configuration
	return config, NoConfigurationSourceFoundError
}

func (l localFileReader) Description() string {
	return l.Desc
}

func (l localFileReader) Location() string {
	return l.Path
}
