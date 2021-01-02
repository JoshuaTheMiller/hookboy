package generators

import (
	"testing"

	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	"github.com/hookboy/source/hookboy/internal"
	_ "github.com/hookboy/source/hookboy/prep/generators/explicit"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

func Test_CustomHook_Name_ReturnsNameAsExpected(t *testing.T) {
	l := customHookGenerator{}

	name := l.Name()

	expectedName := "Custom Hooks"
	actualName := name

	if expectedName != actualName {
		t.Error("Name was not as expected")
	}
}

func Test_explicitHooks_Generate_PropagatesError(t *testing.T) {
	var errorToPropagate = hookboy.NewError("some error")

	c := conf.Configuration{
		Hooks: []conf.Hooks{
			conf.Hooks{
				HookName:  "commit-msg",
				Statement: "doesn't matter",
			},
		},
	}
	l := customHookGenerator{
		initialized: true,
		preparer: testPreparer{
			Error: errorToPropagate,
		},
	}

	ef, ftc, err := l.Generate(c, nil)

	if err != errorToPropagate {
		t.Error("Expected error to be propgated")
	}

	if ef != nil {
		t.Error("Expected ExecutableFiles array to be nil")
	}

	if ftc != nil {
		t.Error("Expected FilesToCreate array to be nil")
	}
}

func Test_CustomHook_Generate_ReturnsAsExpected(t *testing.T) {
	l := customHookGenerator{}

	expectedHook := "commit-msg"
	c := conf.Configuration{
		AutoAddHooks: conf.No,
		Hooks: []conf.Hooks{
			conf.Hooks{
				// Will prepare the hookfile itself
				HookName: expectedHook,
				// Will also prepare another file for this statement
				Statement: "S-1",
				Files: []conf.HookFile{
					conf.HookFile{
						Path: "prepboy_test.go",
					},
				},
			},
		},
	}

	readDirForTest := readDirForTest{}

	ef, ftc, err := l.Generate(c, readDirForTest.readDir)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	// One for the listed hookfile, one for the statement file that will be generated
	expectedLengthOfEf := 2
	actualLengthOfEf := len(ef)
	expectedLengthOfFtc := 1
	actualLengthOfFtc := len(ftc)
	if expectedLengthOfEf != actualLengthOfEf || expectedLengthOfFtc != actualLengthOfFtc {
		t.Errorf("Amount of files returned is incorrect.\n e EF: %d, a EF: %d\n e FTC: %d, a FTC: %d", expectedLengthOfEf, actualLengthOfEf, expectedLengthOfFtc, actualLengthOfFtc)
		return
	}

	if ef[0].AssociatedHook != expectedHook {
		t.Errorf("Expected AssociatedHook to be %s", expectedHook)
	}

	expectedPath := "/commit-msg-statement"
	actualPath := ftc[0].Path
	if expectedPath != actualPath {
		t.Errorf("Expected Path to be %s, received '%s'", expectedPath, actualPath)
	}

	expectedContents := "S-1"
	actualContents := ftc[0].Contents
	if expectedContents != actualContents {
		t.Errorf("Expected Contents to be %s, received '%s'", expectedContents, actualContents)
	}
}

type testPreparer struct {
	Error hookboy.Error
}

func (t testPreparer) Name() string {
	return "test"
}

func (t testPreparer) Prepare(conf.Hooks, conf.Configuration) ([]p.ExecutableFile, []internal.FileToCreate, hookboy.Error) {
	return nil, nil, t.Error
}
