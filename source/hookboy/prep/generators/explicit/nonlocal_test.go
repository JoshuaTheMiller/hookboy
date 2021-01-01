package explicit

import (
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

func Test_nonlocal_Name_IsSetAsExpected(t *testing.T) {
	l := nonlocal{}

	name := l.Name()

	expectedName := "nonlocal"
	actualName := name

	if expectedName != actualName {
		t.Error("Name mismatch")
	}
}

func Test_CustomHook_Generate_NonLocalFileReturnsError(t *testing.T) {
	nl := nonlocal{}

	c := conf.Configuration{}
	hooks := conf.Hooks{
		HookName: "commit-msg",
		Files: []conf.HookFile{
			conf.HookFile{
				Path: "https://example.com/some_file.go",
			},
		},
	}

	_, _, err := nl.Prepare(hooks, c)

	if err != nonLocalFileError {
		t.Error("Expected error")
		return
	}
}

func Test_nonlocal_Prepare_OnlyLocalFilesReturnsEmptyArrays(t *testing.T) {
	nl := nonlocal{}

	c := conf.Configuration{}
	hooks := conf.Hooks{
		HookName: "commit-msg",
		Files: []conf.HookFile{
			conf.HookFile{
				Path: "some_file.go",
			},
			conf.HookFile{
				Path: "some_file2.go",
			},
		},
	}

	ef, ftc, err := nl.Prepare(hooks, c)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	expectedEfLength := 0
	actualEfLength := len(ef)
	if expectedEfLength != actualEfLength {
		t.Error("Expected no files as no statement is present")
	}

	expectedFtcLength := 0
	actualFtcLength := len(ftc)
	if expectedFtcLength != actualFtcLength {
		t.Error("Expected no files as no statement is present")
	}
}
