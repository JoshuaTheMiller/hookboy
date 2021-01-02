package internal

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
)

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
	LocateCurrentConfigurationSource() (CurrentConfigurationSource, hookboy.Error)
	RetrieveCurrentConfiguration() (conf.Configuration, hookboy.Error)
}

// CurrentConfigurationSource is returned from the ./LocateCurrentConfigurationSource function
type CurrentConfigurationSource struct {
	Path        string
	Description string
}
