package aply

import "github.com/hookboy/source/hookboy/internal"

func init() {
	a := applierboy{}

	a.instantiate()

	internal.RegisterApplier(a)
}
