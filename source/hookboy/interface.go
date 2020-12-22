package hookboy

// Application defines the methods available for running the HookBoy
// tool. Typically returned by Builder
type Application interface {
	// Rename to Wrangle and Wrangler
	Install() (string, error)
}

// Builder implementations construct services that adhere to the Application
// interface
type Builder interface {
	Construct(configurationPath string) (Application, error)
}

type bob struct {
	// To be used and set in the future for prioritizing package.json checks
	IsNode bool
}

func (b *bob) Construct(configurationPath string) (Application, error) {
	configPath := configurationPath

	if configPath == "" {
		// TODO: do some checking for default file/locations
		configPath = retrieveConfigPath()
	}

	return getConfiguration(configPath)
}

// GetBuilder retrieves the Builder to be used during construction
// of the tool
func GetBuilder() Builder {
	return &bob{}
}
