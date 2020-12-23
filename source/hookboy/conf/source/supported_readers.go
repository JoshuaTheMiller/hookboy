package source

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/conf/deserialization"
)

var configurationReaders = []configurationReader{
	packageJSONFileReader{
		Desc: "Hookboy will use the package.json file for its configuration source if a hookboy section is found within it.",
	},
	localFolderReader{
		Path: ".hookboy",
		Desc: "Hookboy will look first look for raw hooks in a .hookboy folder, then will look for configuration in a yaml file named .hookboy",
	},
	localFileReader{
		Path:         ".hookboy",
		Desc:         "Hookboy will look first look for raw hooks in a .hookboy folder, then will look for configuration in a yaml file named .hookboy",
		Deserializer: deserialization.YamlDeserializer{},
	},
	localFileReader{
		Path:         ".hookboy.yml",
		Desc:         "Hookboy will look for configuration in a yaml file named .hookboy",
		Deserializer: deserialization.YamlDeserializer{},
	},
}

type configurationReader interface {
	CanRead() bool
	Read() (conf.Configuration, error)
	Description() string
	Location() string
}
