package internal

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
)

var generators []FileGenerator

// RegisterFileGenerator allows for registration of a FileGenerator for further
// use through GetConfigurationExposer()
func RegisterFileGenerator(ftcg FileGenerator) {
	generators = append(generators, ftcg)
}

// GetFileGenerator returns an object that implements ConfigurationExposer
func GetFileGenerator() FileGenerator {
	return generatorComposer{}
}

// A FileGenerator returns files that should be created by later stages
type FileGenerator interface {
	Name() string
	// Generate When called, will generate files if the specific generator deems it appropriate
	// based on the configuration
	Generate(conf.Configuration, func(dirname string) ([]SimpleFile, error)) ([]ExecutableFile, []internal.FileToCreate, error)
}

type generatorComposer struct{}

func (gc generatorComposer) Name() string {
	return "Composer"
}

func (gc generatorComposer) Generate(c conf.Configuration, readDir func(dirname string) ([]SimpleFile, error)) ([]ExecutableFile, []internal.FileToCreate, error) {
	var ef = []ExecutableFile{}
	var ftc = []internal.FileToCreate{}

	for _, generator := range generators {
		var executableFiles, filesToCreate, err = generator.Generate(c, readDir)

		if err != nil {
			return nil, nil, err
		}

		ftc = append(ftc, filesToCreate...)
		ef = append(ef, executableFiles...)
	}

	return ef, ftc, nil
}

// An ExecutableFile is one that has been found that can be executed by a Git Hook
type ExecutableFile struct {
	AssociatedHook string
	Path           string
	ExtraArguments []conf.ExtraArguments
}

// SimpleFile TODO: Does this interface already exist?
type SimpleFile interface {
	Name() string
}
