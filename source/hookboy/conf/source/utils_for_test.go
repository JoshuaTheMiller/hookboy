package source

import (
	"io/ioutil"
	"os"
)

type fileSystemObjectOptions struct {
	Name         string
	FileContents string
	IsFolder     bool
}

func creatFileSystemObjectForTest(options ...fileSystemObjectOptions) {
	for _, opt := range options {
		if opt.IsFolder {
			os.Mkdir(opt.Name, os.ModePerm)
			continue
		}

		ioutil.WriteFile(opt.Name, []byte(opt.FileContents), os.ModePerm)
	}
}

func deleteFileSystemObjectForTest(options ...fileSystemObjectOptions) {
	for _, opt := range options {
		os.Remove(opt.Name)
	}
}
