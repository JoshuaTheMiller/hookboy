package generators

type fileToCreate struct {
	path     string
	contents string
}

func (f fileToCreate) Path() string     { return f.path }
func (f fileToCreate) Contents() string { return f.contents }
