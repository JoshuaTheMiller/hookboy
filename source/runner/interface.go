package runner

// HookBoy defines the methods available for running the HookBoy
// tool.
type HookBoy interface {
	Install() (string, error)
}
