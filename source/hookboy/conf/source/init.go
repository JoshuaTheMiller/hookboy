package source

import "github.com/hookboy/source/hookboy/internal"

func init() {
	internal.RegisterfigurationExposer(configurationExposer{})
}
