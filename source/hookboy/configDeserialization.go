package hookboy

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func getConfiguration(pathToConfig string) (*Configuration, error) {

	yamlFile, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		var errorString = fmt.Sprintf("cannot read file '%s', please check that it is valid", pathToConfig)
		return nil, errors.New(errorString)
	}

	return deserializeConfiguration(yamlFile)
}

func deserializeConfiguration(rawConfiguration []byte) (*Configuration, error) {
	c := &Configuration{}
	err := yaml.Unmarshal(rawConfiguration, c)
	if err != nil {
		return nil, errors.New("failed to parse configuration file")
	}

	return c.setDefaults(), nil
}

// GetDefaultConfiguration use the config from the
// '.gitgrapple.yml' file
func GetDefaultConfiguration() (*Configuration, error) {
	return getConfiguration(retrieveConfigPath())
}

func (c *Configuration) setDefaults() *Configuration {
	if c.LocalHookDir == "" {
		c.LocalHookDir = defaultLocalHooksDir
	}

	if c.AutoAddHooks == "" {
		c.AutoAddHooks = ByFileName
	}

	return c
}
