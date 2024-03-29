package source

import (
	"github.com/hookboy/source/hookboy/conf/deserialization"
)

var configurationReaders = []configurationReader{
	localFileReader{
		Path:         ".hookboy.yml",
		Desc:         "Hookboy will look for configuration in a yaml file named .hookboy",
		Deserializer: deserialization.YamlDeserializer{},
	},
	localFileReader{
		Path:         ".hookboy",
		Desc:         "Hookboy will look first look for raw hooks in a .hookboy folder, then will look for configuration in a yaml file named .hookboy",
		Deserializer: deserialization.YamlDeserializer{},
	},
	packageJSONFileReader{
		Desc: "Hookboy will use the package.json file for its configuration source if a hookboy section is found within it.",
	},
	localFolderReader{
		Path: ".hookboy",
		Desc: "Hookboy will look first look for raw hooks in a .hookboy folder, then will look for configuration in a yaml file named .hookboy",
	},
}
