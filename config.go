package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type extraArguments struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type hookFile struct {
	Path           string           `yaml:"path"`
	ExtraArguments []extraArguments `yaml:"extraArguments"`
}

type hooks struct {
	HookName  string     `yaml:"hookName"`
	Statement string     `yaml:"statement"`
	Files     []hookFile `yaml:"files"`
}

type autoAddHooks string

const (
	no         autoAddHooks = "No"
	byFileName autoAddHooks = "ByFileName"
)

type configuration struct {
	LocalHookDir string `yaml:"localHookDir"`
	// AutoAddHooks Defaults to ByFileName
	AutoAddHooks autoAddHooks `yaml:"autoAddHooks"`
	Hooks        []hooks      `yaml:"hooks"`
}

func getConfiguration(pathToConfig string) configuration {

	yamlFile, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	c := &configuration{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c.setDefaults()
}

func getDefaultConfiguration() configuration {
	return getConfiguration(".gitgrapple.yml")
}

func (c *configuration) setDefaults() configuration {
	if c.LocalHookDir == "" {
		c.LocalHookDir = "./hooks"
	}

	if c.AutoAddHooks == "" {
		c.AutoAddHooks = byFileName
	}

	return *c
}
