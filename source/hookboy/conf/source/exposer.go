package source

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

type configurationExposer struct {
}

// LocateCurrentConfigurationSource can be used to help determine where current configuration is
// coming from.
func (c configurationExposer) LocateCurrentConfigurationSource() (internal.CurrentConfigurationSource, error) {
	for _, reader := range configurationReaders {
		var sourceExists = reader.CanRead()

		if sourceExists {
			return internal.CurrentConfigurationSource{
				Path:        reader.Location(),
				Description: reader.Description(),
			}, nil
		}
	}

	return internal.CurrentConfigurationSource{}, internal.NoConfigurationSourceFoundError
}

// RetrieveCurrentConfiguration retrieves the current configuration, or returns an
// error if no source of Configuration can be found or if there are issues with consuming
// the configuration.
func (c configurationExposer) RetrieveCurrentConfiguration() (conf.Configuration, error) {
	for _, reader := range configurationReaders {
		var sourceExists = reader.CanRead()

		if sourceExists {
			return reader.Read()
		}
	}

	return conf.Configuration{}, internal.NoConfigurationSourceFoundError
}
