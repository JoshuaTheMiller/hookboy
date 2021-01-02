package boundary

import "os"

type FileWriter func(filename string, data []byte, perm os.FileMode) error
type FolderCreator func(path string, perm os.FileMode) error
