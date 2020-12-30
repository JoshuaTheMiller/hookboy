package internal

import "github.com/hookboy/source/hookboy/conf"

// Prepper is used to prepare files to usage by Git Hook
type Prepper interface {
	PrepareHookfileInfo(c conf.Configuration) (ftc []FileToCreate, e error)
}

var registeredPrepper Prepper

func RegisterPrepper(p Prepper) {
	registeredPrepper = p
}

func GetPrepper() Prepper {
	return registeredPrepper
}

type PrepError struct {
	Description   string
	InternalError error
}

func (pe PrepError) Error() string {
	return pe.Description
}
