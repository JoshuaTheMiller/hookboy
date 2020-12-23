package source

import "github.com/hookboy/source/hookboy/conf"

var configurationSources = []configurationSource{
	configurationSource{
		Reader: packageJSONFileReader{
			Desc: "Hookboy will use the package.json file for its configuration source if a hookboy section is found within it.",
		},
	},
	configurationSource{
		Reader: localFolderReader{
			Path: ".hookboy",
			Desc: "Hookboy will look first look for raw hooks in a .hookboy folder, then will look for configuration in a yaml file named .hookboy",
		},
	},
	configurationSource{
		Reader: localFileReader{
			Path: ".hookboy",
			Desc: "Hookboy will look first look for raw hooks in a .hookboy folder, then will look for configuration in a yaml file named .hookboy",
		},
	},
	configurationSource{
		Reader: localFileReader{
			Path: ".hookboy.yml",
			Desc: "Hookboy will look for configuration in a yaml file named .hookboy",
		},
	},
}

type configurationSource struct {
	Reader configurationReader
}

type configurationReader interface {
	CanRead() bool
	Read() (conf.Configuration, error)
	Description() string
	Location() string
}
