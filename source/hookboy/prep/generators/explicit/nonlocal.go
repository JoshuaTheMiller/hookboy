package explicit

import (
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type nonlocal struct {
}

func (nl nonlocal) Name() string {
	return "nonlocal"
}

func (nl nonlocal) Prepare(h conf.Hooks, c conf.Configuration) (ef []p.ExecutableFile, ftc []internal.FileToCreate, e error) {
	for _, fileToInclude := range h.Files {
		var path = fileToInclude.Path

		if !pathIsNonLocalPath(path) {
			// This func only handles non-local files, skip if not non-local
			continue
		}

		return nil, nil, nonLocalFileError
	}

	return ef, ftc, nil
}

var nonLocalFileError = internal.PrepError{
	Description: "Non-local files are not yet supported!",
}
