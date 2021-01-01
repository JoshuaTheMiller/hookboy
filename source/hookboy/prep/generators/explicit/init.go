package explicit

import "github.com/hookboy/source/hookboy/prep/generators/internal"

func init() {
	internal.Register(local{})
	internal.Register(nonlocal{})
	internal.Register(statement{})
}
