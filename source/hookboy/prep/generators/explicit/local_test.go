package explicit

import (
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

func Test_local_Name_IsSetAsExpected(t *testing.T) {
	l := local{}

	name := l.Name()

	expectedName := "local"
	actualName := name

	if expectedName != actualName {
		t.Error("Name mismatch")
	}
}

func Test_local_Prepare_OnlyNonLocalFilesReturnsEmptyArrays(t *testing.T) {
	l := local{}

	c := conf.Configuration{}
	hooks := conf.Hooks{
		HookName: "commit-msg",
		Files: []conf.HookFile{
			conf.HookFile{
				Path: "https://example.com/some_file.go",
			},
			conf.HookFile{
				Path: "https://example.com/some_file2.go",
			},
		},
	}

	ef, ftc, err := l.Prepare(hooks, c)

	if err != nil {
		t.Error("Expected error to be nil")
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

func Test_local_Prepare_OnlyReturnsExecutableFilesAsFilesAlreadyExistLocally(t *testing.T) {
	l := local{}

	hookname := "commit-msg"
	localFilePath := "/some_file.go"
	c := conf.Configuration{}
	hooks := conf.Hooks{
		HookName: hookname,
		Files: []conf.HookFile{
			conf.HookFile{
				Path: localFilePath,
			},
		},
	}

	ef, ftc, err := l.Prepare(hooks, c)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	expectedEfLength := 1
	actualEfLength := len(ef)
	if expectedEfLength != actualEfLength {
		t.Error("Expected no files as no statement is present")
	}

	if ef[0].AssociatedHook != hookname {
		t.Error("Incorrect hookname")
	}

	if ef[0].Path != localFilePath {
		t.Error("Incorrect hookname")
	}

	expectedFtcLength := 0
	actualFtcLength := len(ftc)
	if expectedFtcLength != actualFtcLength {
		t.Error("Expected no files as no statement is present")
	}
}
