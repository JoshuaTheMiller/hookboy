package deserialization

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
)

// Deserializer objects take raw bytes, and return configuration
type Deserializer interface {
	Deserialize(raw []byte) (conf.Configuration, hookboy.Error)
}
