package hookboy

import (
	"github.com/hookboy/source/hookboy/aply"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/conf/source"
)

// Application defines the methods available for running the HookBoy
// tool. Typically returned by Builder
type Application interface {
	// Rename to Wrangle and Wrangler
	Install() (string, error)
	CurrentConfiguration() (conf.Configuration, error)
	ConfigurationLocation() (string, error)
}

// Builder implementations construct services that adhere to the Application
// interface
type Builder interface {
	Construct() (Application, error)
}

type hookboyTheAppliction struct {
	Configuration conf.Configuration
}

func (hb *hookboyTheAppliction) Install() (string, error) {
	return aply.Install(hb.Configuration)
}

func (hb *hookboyTheAppliction) CurrentConfiguration() (conf.Configuration, error) {
	return hb.Configuration, nil
}

func (hb *hookboyTheAppliction) ConfigurationLocation() (string, error) {
	var source, err = source.LocateCurrentConfigurationSource()
	return source.Path, err
}

type bob struct {
}

func (b *bob) Construct() (Application, error) {
	configuration, err := source.RetrieveCurrentConfiguration()

	if err != nil {
		return nil, err
	}

	return &hookboyTheAppliction{
		Configuration: configuration,
	}, nil
}

// GetBuilder retrieves the Builder to be used during construction
// of the tool
func GetBuilder() Builder {
	return &bob{}
}
