package deserialization

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"gopkg.in/yaml.v3"
)

// YamlDeserializer useful for deserializing yaml files
type YamlDeserializer struct {
}

// Deserialize given raw bytes that represent yaml, this will return the configuration
func (d YamlDeserializer) Deserialize(rawConfiguration []byte) (conf.Configuration, hookboy.Error) {
	c := &conf.Configuration{}
	err := yaml.Unmarshal(rawConfiguration, c)
	if err != nil {
		return conf.Configuration{}, hookboy.WrapError(err, "failed to parse configuration file")
	}

	return *c.SetDefaults(), nil
}
