package deserialization

import (
	"errors"

	"github.com/hookboy/source/hookboy/conf"
	"gopkg.in/yaml.v3"
)

// YamlDeserializer useful for deserializing yaml files
type YamlDeserializer struct {
}

// Deserialize given raw bytes that represent yaml, this will return the configuration
func (d YamlDeserializer) Deserialize(rawConfiguration []byte) (conf.Configuration, error) {
	c := &conf.Configuration{}
	err := yaml.Unmarshal(rawConfiguration, c)
	if err != nil {
		return conf.Configuration{}, errors.New("failed to parse configuration file")
	}

	return *c.SetDefaults(), nil
}
