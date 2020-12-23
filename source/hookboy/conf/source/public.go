package source

import "github.com/hookboy/source/hookboy/conf"

// CurrentConfigurationSource is returned from the ./LocateCurrentConfigurationSource function
type CurrentConfigurationSource struct {
	Path        string
	Description string
}

// LocateCurrentConfigurationSource can be used to help determine where current configuration is
// coming from.
func LocateCurrentConfigurationSource() (CurrentConfigurationSource, error) {
	for _, reader := range configurationReaders {
		var sourceExists = reader.CanRead()

		if sourceExists {
			return CurrentConfigurationSource{
				Path:        reader.Location(),
				Description: reader.Description(),
			}, nil
		}
	}

	return CurrentConfigurationSource{}, NoConfigurationSourceFoundError
}

// RetrieveCurrentConfiguration retrieves the current configuration, or returns an
// error if no source of Configuration can be found or if there are issues with consuming
// the configuration.
func RetrieveCurrentConfiguration() (conf.Configuration, error) {
	for _, reader := range configurationReaders {
		var sourceExists = reader.CanRead()

		if sourceExists {
			return reader.Read()
		}
	}

	return conf.Configuration{}, NoConfigurationSourceFoundError
}

// NoConfigurationSourceFoundError returned when no source of configuration can be found.
var NoConfigurationSourceFoundError = ConfigurationSourceError{
	Description: "No source of configuration found.",
}

// ConfigurationSourceError is the generic error that is returned if there is any form of configuration
// errors
type ConfigurationSourceError struct {
	Description string
}

func (e ConfigurationSourceError) Error() string {
	return e.Description
}
