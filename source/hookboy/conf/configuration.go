package conf

// ExtraArguments that may be passed to git hooks at run time
// These pairs will be hardcoded into the final hooks file present
// in the .git/hooks directory
type ExtraArguments struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// HookFile represents the information necessary to execute a file
// during the running of a git hook
type HookFile struct {
	Path           string           `yaml:"path"`
	ExtraArguments []ExtraArguments `yaml:"extraArguments"`
}

// Hooks are the various different git hooks that will be
// generated during installation
type Hooks struct {
	HookName  string     `yaml:"hookName"`
	Statement string     `yaml:"statement"`
	Files     []HookFile `yaml:"files"`
}

// AutoAddHooks the different recognized options available
// for auto adding hooks
type AutoAddHooks string

const (
	// No do not auto add any hooks
	No AutoAddHooks = "No"
	// ByFileName when set, will cause all files that match
	// existing git hook names to be used during installation
	ByFileName AutoAddHooks = "ByFileName"
)

// Configuration contains the various values that are necessary
// for installing git hooks using this application
type Configuration struct {
	LocalHookDir string `yaml:"localHookDir"`
	// AutoAddHooks Defaults to ByFileName
	AutoAddHooks AutoAddHooks `yaml:"autoAddHooks"`
	// CacheDirectory defaults to '.git/hooks/.hookboy-cache'
	CacheDirectory string  `yaml:"cacheDirectory"`
	Hooks          []Hooks `yaml:"hooks"`
}

// SetDefaults sets appropriate defaults
func (c *Configuration) SetDefaults() *Configuration {
	if c.LocalHookDir == "" {
		c.LocalHookDir = DefaultLocalHooksDir
	}

	if c.AutoAddHooks == "" {
		c.AutoAddHooks = ByFileName
	}

	if c.CacheDirectory == "" {
		c.CacheDirectory = ".git/hooks/.hookboy-cache"
	}

	return c
}

// GetCacheDirectory returns the cache directory dictacted by the configuration
func (c Configuration) GetCacheDirectory() string {
	return c.CacheDirectory
}
