package generators

import "github.com/hookboy/source/hookboy/prep/internal"

func init() {
	internal.RegisterFileGenerator(customHookGenerator{})
	internal.RegisterFileGenerator(localHooksGenerator{})
}
