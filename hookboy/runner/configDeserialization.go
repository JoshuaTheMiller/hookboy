package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func getConfiguration(pathToConfig string) (*configuration, error) {

	yamlFile, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)

		var errorString = fmt.Sprintf("cannot read file '%s', please check that it is valid", pathToConfig)
		return nil, errors.New(errorString)
	}

	return deserializeConfiguration(yamlFile)
}

func deserializeConfiguration(rawConfiguration []byte) (*configuration, error) {
	c := &configuration{}
	err := yaml.Unmarshal(rawConfiguration, c)
	if err != nil {
		//log.Fatalf("Unmarshal: %v", err)

		return nil, errors.New("failed to parse configuration file")
	}

	return c.setDefaults(), nil
}

func getDefaultConfiguration() (*configuration, error) {
	return getConfiguration(".gitgrapple.yml")
}

func (c *configuration) setDefaults() *configuration {
	if c.LocalHookDir == "" {
		c.LocalHookDir = "./hooks"
	}

	if c.AutoAddHooks == "" {
		c.AutoAddHooks = byFileName
	}

	return c
}
