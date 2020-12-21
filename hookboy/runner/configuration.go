package runner

// ExtraArguments that may be passed to git hooks at run time
// These pairs will be hardcoded into the final hooks file present
// in the .git/hooks directory
type extraArguments struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// HookFile represents the information necessary to execute a file
// during the running of a git hook
type hookFile struct {
	Path           string           `yaml:"path"`
	ExtraArguments []extraArguments `yaml:"extraArguments"`
}

// Hooks are the various different git hooks that will be
// generated during installation
type hooks struct {
	HookName  string     `yaml:"hookName"`
	Statement string     `yaml:"statement"`
	Files     []hookFile `yaml:"files"`
}

// AutoAddHooks the different recognized options available
// for auto adding hooks
type autoAddHooks string

const (
	// No do not auto add any hooks
	no autoAddHooks = "No"
	// ByFileName when set, will cause all files that match
	// existing git hook names to be used during installation
	byFileName autoAddHooks = "ByFileName"
)

// Configuration contains the various values that are necessary
// for installing git hooks using this application
type configuration struct {
	LocalHookDir string `yaml:"localHookDir"`
	// AutoAddHooks Defaults to ByFileName
	AutoAddHooks autoAddHooks `yaml:"autoAddHooks"`
	Hooks        []hooks      `yaml:"hooks"`
}
