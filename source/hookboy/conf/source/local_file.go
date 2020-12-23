package source

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/conf/deserialization"
)

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
	rawFile, err := readFile(l.Path)

	if err != nil {
		return conf.Configuration{}, err
	}

	var des = deserialization.YamlDeserializer{}

	configuration, err := des.Deserialize(rawFile)

	if err != nil {
		return conf.Configuration{}, err
	}

	return configuration, nil
}

func (l localFileReader) Description() string {
	return l.Desc
}

func (l localFileReader) Location() string {
	return l.Path
}
