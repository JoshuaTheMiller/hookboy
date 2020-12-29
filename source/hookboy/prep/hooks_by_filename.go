package prep

import (
	"github.com/hookboy/source/hookboy/conf"
)

func getHooksByFileName(localHooksDir string, readDir func(dirname string) ([]simpleFile, error)) (eFiles []executableFile, err error) {
	var files, readErr = readDir(localHooksDir)

	if readErr != nil {
		return nil, readErr
	}

	for _, f := range files {
		var potentialHookName = f.Name()

		if fileIsARecognizedHook(potentialHookName) {
			eFiles = append(eFiles, executableFile{
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
