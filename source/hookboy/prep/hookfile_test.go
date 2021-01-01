package prep

import (
	"testing"

	"github.com/hookboy/source/hookboy/conf"
	_ "github.com/hookboy/source/hookboy/prep/generators/explicit"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

var fileAExpectedContents = `#!/bin/sh
retVal0=exec "A_PATH_1" "$@" EA_N_1=EA_V_1
retVal0=$?
retVal1=exec "A_PATH_2" "$@"
retVal1=$?


if [ $retVal0 -ne 0 ] || [ $retVal1 -ne 0 ];
then
exit 1
fi
exit 0`

func TestGenerateHookFileContents(t *testing.T) {
	ef := []p.ExecutableFile{
		p.ExecutableFile{
			AssociatedHook: "A",
			Path:           "A_PATH_1",
			ExtraArguments: []conf.ExtraArguments{
				conf.ExtraArguments{
					Name:  "EA_N_1",
					Value: "EA_V_1",
				},
			},
		},
		p.ExecutableFile{
			AssociatedHook: "A",
			Path:           "A_PATH_2",
		},
		p.ExecutableFile{
			AssociatedHook: "B",
			Path:           "B_PATH_2",
		},
	}

	hookFileContents := generateHookFileContents(ef, conf.Configuration{})

	amountOfContentsGenerated := len(hookFileContents)
	if amountOfContentsGenerated != 2 {
		t.Errorf("Expected the contents of 2 hookfiles to be created, found %d", amountOfContentsGenerated)
	}

	// var fileA = hookFileContents[0]

	// if fileA.Contents() != fileAExpectedContents {
	// 	t.Errorf("Expected contents to be\n'%s'\n, but found \n'%s'\n", fileAExpectedContents, fileA.Contents())
	// }
}
