package source

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hookboy/source/hookboy"
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

func (l packageJSONFileReader) Read() (conf.Configuration, hookboy.Error) {

	var packageJSONPath = packageJSONName

	result := checkForFileSystemObject(packageJSONPath)

	if result == doesNotExist || result == folder {
		return conf.Configuration{}, hookboy.NewError("Error reading package.json")
	}

	data, err := ioutil.ReadFile(packageJSONPath)
	if err != nil {
		return conf.Configuration{}, hookboy.NewError("Error reading package.json")
	}

	type packageJSON struct {
		Hookboy *conf.Configuration `json:"hookboy"`
	}

	var obj packageJSON

	err = json.Unmarshal(data, &obj)

	if err != nil {
		return conf.Configuration{}, hookboy.NewError("Error reading package.json")
	}

	if obj.Hookboy != nil {
		return *obj.Hookboy, nil
	}

	return conf.Configuration{}, hookboy.NewError("Error reading package.json")
}

func (l packageJSONFileReader) Description() string {
	return l.Desc
}

func (l packageJSONFileReader) Location() string {
	return packageJSONName
}
