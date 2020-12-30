package prep

import "github.com/hookboy/source/hookboy/internal"

func init() {
	p := prepboy{}

	p.instantiate()

	internal.RegisterPrepper(p)
}
