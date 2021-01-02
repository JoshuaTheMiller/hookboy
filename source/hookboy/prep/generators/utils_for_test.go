package generators

import (
	"github.com/hookboy/source/hookboy"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type readDirForTest struct {
	Files []p.SimpleFile
	Error hookboy.Error
}

func (rdft readDirForTest) readDir(dirname string) ([]p.SimpleFile, hookboy.Error) {
	return rdft.Files, rdft.Error
}
