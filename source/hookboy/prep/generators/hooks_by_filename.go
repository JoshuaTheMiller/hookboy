package generators

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type localHooksGenerator struct {
}

func (l localHooksGenerator) Generate(c conf.Configuration, readDir func(dirname string) ([]p.SimpleFile, hookboy.Error)) (ef []p.ExecutableFile, ftc []internal.FileToCreate, err hookboy.Error) {
	if c.AutoAddHooks != conf.ByFileName {
		return []p.ExecutableFile{}, []internal.FileToCreate{}, nil
	}

	files, err := getHooksByFileName(c.LocalHookDir, readDir)

	if err != nil {
		return nil, nil, hookboy.WrapError(err, "Error prepping hooks by filename")
	}

	return files, []internal.FileToCreate{}, nil
}

func (l localHooksGenerator) Name() string {
	return "Local Hook Files"
}

func getHooksByFileName(localHooksDir string, readDir func(dirname string) ([]p.SimpleFile, hookboy.Error)) (eFiles []p.ExecutableFile, err hookboy.Error) {
	var files, readErr = readDir(localHooksDir)

	if readErr != nil {
		return nil, readErr
	}

	for _, f := range files {
		var potentialHookName = f.Name()

		if fileIsARecognizedHook(potentialHookName) {
			eFiles = append(eFiles, p.ExecutableFile{
				AssociatedHook: potentialHookName,
				Path:           "./" + localHooksDir + "/" + potentialHookName,
			})
		}
	}

	return eFiles, nil
}

func fileIsARecognizedHook(fileName string) bool {
	return itemExists(conf.RecognizedHooks, fileName)
}
