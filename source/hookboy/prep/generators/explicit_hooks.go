package generators

import (
	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	g "github.com/hookboy/source/hookboy/prep/generators/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

type customHookGenerator struct {
	initialized bool
	preparer    g.ExplicitHookPreparer
}

func (chg customHookGenerator) Generate(c conf.Configuration, readDir func(dirname string) ([]p.SimpleFile, hookboy.Error)) (ef []p.ExecutableFile, ftc []internal.FileToCreate, err hookboy.Error) {
	chg.initialize()
	var preparer = chg.preparer

	for _, hook := range c.Hooks {
		var newExecutableFiles, newFilesToCreate, newPrepareError = preparer.Prepare(hook, c)

		if newPrepareError != nil {
			return nil, nil, newPrepareError
		}

		ef = append(ef, newExecutableFiles...)
		ftc = append(ftc, newFilesToCreate...)
	}

	return ef, ftc, err
}

func (chg *customHookGenerator) initialize() {
	if !chg.initialized {
		chg.preparer = g.GetExplicitHooksPreparer()
		chg.initialized = true
	}
}

func (chg customHookGenerator) Name() string {
	return "Custom Hooks"
}
