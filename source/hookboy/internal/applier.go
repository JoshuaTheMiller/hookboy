package internal

import "github.com/hookboy/source/hookboy/conf"

// Applier configures hooks
type Applier interface {
	Install(conf.Configuration, []FileToCreate) (string, error)
}

var registeredApplier Applier

// RegisterApplier allows for registration of an Applier for further
// use through GetApplier()
func RegisterApplier(a Applier) {
	registeredApplier = a
}

// GetApplier allows for further use of an Applier registered through
// RegisterApplier
func GetApplier() Applier {
	return registeredApplier
}

// An AplyError is an error generated during the Apply stage
type AplyError struct {
	Description   string
	InternalError error
}

func (ae AplyError) Error() string {
	return ae.Description
}
