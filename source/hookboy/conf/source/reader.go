package source

import "github.com/hookboy/source/hookboy/conf"

type configurationReader interface {
	CanRead() bool
	Read() (conf.Configuration, error)
	Description() string
	Location() string
}
