package source

import "github.com/hookboy/source/hookboy"

// NoConfigurationSourceFoundError returned when no source of configuration can be found.
var NoConfigurationSourceFoundError = hookboy.NewError("No source of configuration found.")
