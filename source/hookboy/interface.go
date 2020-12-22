package hookboy

// Application defines the methods available for running the HookBoy
// tool. Typically returned by Builder
type Application interface {
	Install() (string, error)
}
