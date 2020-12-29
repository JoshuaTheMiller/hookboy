package prep

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
		os.RemoveAll(opt.Name)
	}
}

type simpleFileForTest struct {
	name string
}

func (s simpleFileForTest) Name() string {
	return s.name
}
