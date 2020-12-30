package internal

// FileToCreate contains information about files that should
// be created in later stages
type FileToCreate struct {
	Path     string
	Contents string
}
