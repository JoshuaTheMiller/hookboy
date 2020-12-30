package internal

import "github.com/hookboy/source/hookboy/conf"

var registeredConfigurationExposer ConfigurationExposer

// RegisterfigurationExposer allows for registration of an Exposer for further
// use through GetConfigurationExposer()
func RegisterfigurationExposer(ce ConfigurationExposer) {
	registeredConfigurationExposer = ce
}

// GetConfigurationExposer returns an object that implements ConfigurationExposer
func GetConfigurationExposer() ConfigurationExposer {
	return registeredConfigurationExposer
}

// ConfigurationExposer exposes functions to help gleem more information about configuration
type ConfigurationExposer interface {
	LocateCurrentConfigurationSource() (CurrentConfigurationSource, error)
	RetrieveCurrentConfiguration() (conf.Configuration, error)
}

// NoConfigurationSourceFoundError returned when no source of configuration can be found.
var NoConfigurationSourceFoundError = ConfigurationSourceError{
	Description: "No source of configuration found.",
}

// CurrentConfigurationSource is returned from the ./LocateCurrentConfigurationSource function
type CurrentConfigurationSource struct {
	Path        string
	Description string
}

// ConfigurationSourceError is the generic error that is returned if there is any form of configuration
// errors
type ConfigurationSourceError struct {
	Description string
}

func (e ConfigurationSourceError) Error() string {
	return e.Description
}
