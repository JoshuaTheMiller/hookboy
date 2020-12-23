package deserialization

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hookboy/source/hookboy/conf"
	"gopkg.in/yaml.v3"
)

// GetConfiguration retrieves a deserialized configuration object from
// the given path
func GetConfiguration(pathToConfig string) (*conf.Configuration, error) {

	yamlFile, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		var errorString = fmt.Sprintf("cannot read file '%s', please check that it is valid", pathToConfig)
		return nil, errors.New(errorString)
	}

	return deserializeConfiguration(yamlFile)
}

func deserializeConfiguration(rawConfiguration []byte) (*conf.Configuration, error) {
	c := &conf.Configuration{}
	err := yaml.Unmarshal(rawConfiguration, c)
	if err != nil {
		return nil, errors.New("failed to parse configuration file")
	}

	return c.SetDefaults(), nil
}

// GetDefaultConfiguration use the config from the
// '.gitgrapple.yml' file
func GetDefaultConfiguration() (*conf.Configuration, error) {
	return GetConfiguration(conf.RetrieveConfigPath())
}
