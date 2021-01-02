package source

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
)

type configurationReader interface {
	CanRead() bool
	Read() (conf.Configuration, hookboy.Error)
	Description() string
	Location() string
}
