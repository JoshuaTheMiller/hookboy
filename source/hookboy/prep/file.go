package prep

import "github.com/hookboy/source/hookboy/conf"

// FileToCreate contains information about files that should
// be created in later stages
type FileToCreate interface {
	Path() string
	Contents() string
}

type fileToCreate struct {
	path     string
	contents string
}

func (f fileToCreate) Path() string     { return f.path }
func (f fileToCreate) Contents() string { return f.contents }

type executableFile struct {
	AssociatedHook string
	Path           string
	ExtraArguments []conf.ExtraArguments
}
