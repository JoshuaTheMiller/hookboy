package internal

import (
	"fmt"

	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

// An ExplicitHookPreparer is used to generate files for hooks that are explicitly defined
type ExplicitHookPreparer interface {
	Name() string
	Prepare(conf.Hooks, conf.Configuration) ([]p.ExecutableFile, []internal.FileToCreate, error)
}

// Register registers an additional ExplicitHookPreparer.
// Registering a hook with the same name more than once will cause a Panic!
func Register(ehp ExplicitHookPreparer) {
	panicIfPreparerWithNameIsAlreadyRegistered(ehp)

	preparers = append(preparers, ehp)
}

// GetExplicitHooksPreparer returns the registered ExplicitHookPreparer
func GetExplicitHooksPreparer() ExplicitHookPreparer {
	return compositeExplicitHookPreparer{}
}

func panicIfPreparerWithNameIsAlreadyRegistered(ehp ExplicitHookPreparer) {
	for _, p := range preparers {
		if p.Name() == ehp.Name() {
			var panicMessage = fmt.Sprintf("Preparer with name '%s' has already been registered!", ehp.Name())
			panic(panicMessage)
		}
	}
}

var preparers []ExplicitHookPreparer

type compositeExplicitHookPreparer struct {
}

func (cehp compositeExplicitHookPreparer) Name() string {
	return "composite"
}

func (cehp compositeExplicitHookPreparer) Prepare(h conf.Hooks, c conf.Configuration) (ef []p.ExecutableFile, ftc []internal.FileToCreate, e error) {
	for _, p := range preparers {
		var executableFiles, filesToCreate, err = p.Prepare(h, c)

		if err != nil {
			return nil, nil, err
		}

		ef = append(ef, executableFiles...)
		ftc = append(ftc, filesToCreate...)
	}

	return ef, ftc, nil
}
