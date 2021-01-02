package prep

import (
	"io/ioutil"

	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"

	p "github.com/hookboy/source/hookboy/prep/internal"
)

func (pb prepboy) PrepareHookfileInfo(c conf.Configuration) (ftc []internal.FileToCreate, e hookboy.Error) {
	if pb.Instantiated != true {
		pb.instantiate()
	}

	var fg = p.GetFileGenerator()

	ef, ftc, err := fg.Generate(c, pb.ReadDir)

	if err != nil {
		return nil, err
	}

	var hookfilesToCreate = generateHookFileContents(ef, c)

	ftc = append(ftc, hookfilesToCreate...)

	return ftc, nil
}

type prepboy struct {
	Instantiated bool
	ReadDir      func(dirname string) ([]p.SimpleFile, hookboy.Error)
}

func (pb *prepboy) instantiate() {
	pb.ReadDir = readDir
	pb.Instantiated = true
}

func readDir(dir string) ([]p.SimpleFile, hookboy.Error) {
	var files, err = ioutil.ReadDir(dir)

	if err != nil {
		return nil, hookboy.WrapError(err, "Failed to read directory!")
	}

	data := make([]p.SimpleFile, len(files))

	for i := range files {
		data[i] = p.SimpleFile(files[i])
	}

	return data, nil
}
