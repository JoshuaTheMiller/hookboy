package internal

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
)

// Prepper is used to prepare files to usage by Git Hook
type Prepper interface {
	PrepareHookfileInfo(c conf.Configuration) (ftc []FileToCreate, e hookboy.Error)
}

var registeredPrepper Prepper

func RegisterPrepper(p Prepper) {
	registeredPrepper = p
}

func GetPrepper() Prepper {
	return registeredPrepper
}
