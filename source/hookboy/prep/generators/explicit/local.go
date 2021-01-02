package explicit

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type local struct {
}

func (l local) Name() string {
	return "local"
}

func (l local) Prepare(h conf.Hooks, c conf.Configuration) (ef []p.ExecutableFile, ftc []internal.FileToCreate, e hookboy.Error) {
	for _, fileToInclude := range h.Files {
		var path = fileToInclude.Path

		if pathIsNonLocalPath(path) {
			// This func only handles local files, skip if non-local
			continue
		}

		var file = p.ExecutableFile{
			AssociatedHook: h.HookName,
			Path:           fileToInclude.Path,
			ExtraArguments: fileToInclude.ExtraArguments,
		}

		ef = append(ef, file)
	}

	return ef, ftc, nil
}
