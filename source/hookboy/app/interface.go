package app

import (
	"fmt"

	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

// Application defines the methods available for running the HookBoy
// tool. Typically returned by Builder
type Application interface {
	// Rename to Wrangle and Wrangler
	Install() (string, hookboy.Error)
	CurrentConfiguration() (conf.Configuration, hookboy.Error)
	ConfigurationLocation() (string, hookboy.Error)
}

// Builder implementations construct services that adhere to the Application
// interface
type Builder interface {
	Construct() (Application, hookboy.Error)
}

type hookboyTheAppliction struct {
	Applier       internal.Applier
	Configuration conf.Configuration
	CE            internal.ConfigurationExposer
}

func (hb *hookboyTheAppliction) Install() (string, hookboy.Error) {

	var prepboy = internal.GetPrepper()

	var filesToCreate, err = prepboy.PrepareHookfileInfo(hb.Configuration)

	if err != nil {
		return "", err
	}

	return hb.Applier.Install(hb.Configuration, filesToCreate)
}

func (hb *hookboyTheAppliction) CurrentConfiguration() (conf.Configuration, hookboy.Error) {
	return hb.Configuration, nil
}

func (hb *hookboyTheAppliction) ConfigurationLocation() (string, hookboy.Error) {
	var source, err = hb.CE.LocateCurrentConfigurationSource()
	var message = fmt.Sprintf("| Configuration Source\n|--> Source: %s\n|--> Description: %s", source.Path, source.Description)
	return message, err
}

type bob struct {
}

func (b *bob) Construct() (Application, hookboy.Error) {
	ce := internal.GetConfigurationExposer()

	configuration, err := ce.RetrieveCurrentConfiguration()

	if err != nil {
		return nil, err
	}

	return &hookboyTheAppliction{
		Configuration: configuration,
		CE:            ce,
		Applier:       internal.GetApplier(),
	}, nil
}

// GetBuilder retrieves the Builder to be used during construction
// of the tool
func GetBuilder() Builder {
	return &bob{}
}
