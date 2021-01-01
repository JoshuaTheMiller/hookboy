package explicit

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	"github.com/hookboy/source/hookboy/internal/util"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type nonlocal struct {
	Initialized bool
	HTTPGetter  util.HTTPGetter
}

func (nl nonlocal) Name() string {
	return "nonlocal"
}

func (nl nonlocal) Prepare(h conf.Hooks, c conf.Configuration) (ef []p.ExecutableFile, ftc []internal.FileToCreate, e error) {
	nl.initialize()

	for index, fileToInclude := range h.Files {
		var path = fileToInclude.Path

		if !pathIsNonLocalPath(path) {
			// This func only handles non-local files, skip if not non-local
			continue
		}

		var contents, err = nl.HTTPGetter.Get(path)

		if err != nil {
			var fileError = fmt.Sprintf("Failed to retrieve non-local file. Please validate that you have access to the file, and that the configured URL is correct: %s", path)
			return nil, nil, internal.PrepError{
				Description:   fileError,
				InternalError: err,
			}
		}

		var localPath = generateLocalPath(path, c, index)

		var fileToCreate = internal.FileToCreate{
			Contents: contents,
			Path:     localPath,
		}

		var executableFile = p.ExecutableFile{
			AssociatedHook: h.HookName,
			Path:           localPath,
		}

		return []p.ExecutableFile{executableFile}, []internal.FileToCreate{fileToCreate}, nil
	}

	return ef, ftc, nil
}

func generateLocalPath(path string, c conf.Configuration, index int) string {
	var cachDir = c.GetCacheDirectory()

	// We can trust that path is a valid URL at this point
	u, _ := url.Parse(path)
	uParts := strings.Split(u.Host, ".")

	return fmt.Sprintf("%s/%s-%d", cachDir, uParts[0], index)
}

func (nl *nonlocal) initialize() {
	if nl.Initialized == false {
		nl.Initialized = true
		nl.HTTPGetter = util.DefaultHTTPGetter()
	}
}

var nonLocalFileError = internal.PrepError{
	Description: "Non-local files are not yet supported!",
}
