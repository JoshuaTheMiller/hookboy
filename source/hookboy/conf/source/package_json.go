package source

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hookboy/source/hookboy/conf"
)

const packageJSONName = "package.json"

type packageJSONFileReader struct {
	Desc string
}

func (l packageJSONFileReader) CanRead() bool {
	var packageJSONPath = packageJSONName

	result := checkForFileSystemObject(packageJSONPath)

	if result == doesNotExist || result == folder {
		return false
	}

	data, err := ioutil.ReadFile(packageJSONPath)
	if err != nil {
		return false
	}

	type packageJSON struct {
		Hookboy interface{} `json:"hookboy"`
	}

	var obj packageJSON

	err = json.Unmarshal(data, &obj)

	if err != nil {
		return false
	}

	if obj.Hookboy != nil {
		return true
	}

	return false
}

func (l packageJSONFileReader) Read() (conf.Configuration, error) {
	var config conf.Configuration
	return config, NoConfigurationSourceFoundError
}

func (l packageJSONFileReader) Description() string {
	return l.Desc
}

func (l packageJSONFileReader) Location() string {
	return packageJSONName
}
