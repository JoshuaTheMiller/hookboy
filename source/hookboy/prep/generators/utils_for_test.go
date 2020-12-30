package generators

import p "github.com/hookboy/source/hookboy/prep/internal"

type readDirForTest struct {
	Files []p.SimpleFile
	Error error
}

func (rdft readDirForTest) readDir(dirname string) ([]p.SimpleFile, error) {
	return rdft.Files, rdft.Error
}
