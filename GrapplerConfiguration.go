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

type configuration struct {
	LocalHookDir                      string            `yaml:"localHookDir"`
	DoNotAutoAddHooksFromLocalHookDir bool              `yaml:"doNotAutoAddHooksFromLocalHookDir"`
	Hooks                             []hooks           `yaml:"hooks"`
	HookStatements                    map[string]string `yaml:"hookStatements"`
}

func getConfiguration() *configuration {

	yamlFile, err := ioutil.ReadFile(".gitgrapple.yml")
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

func (c *configuration) setDefaults() *configuration {
	if c.LocalHookDir == "" {
		c.LocalHookDir = "./hooks"
	}

	return c
}
