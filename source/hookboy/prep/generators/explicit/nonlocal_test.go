package explicit

import (
	"fmt"
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

func Test_CustomHook_Generate_NonLocalFileReturnsExpectedNonLocalFile(t *testing.T) {
	nl := nonlocal{}

	cachDir := "somecachedir"
	c := conf.Configuration{
		CacheDirectory: cachDir,
	}
	hooks := conf.Hooks{
		HookName: "commit-msg",
		Files: []conf.HookFile{
			conf.HookFile{
				Path: "https://example.com/some_file.go",
			},
		},
	}

	ef, ftc, err := nl.Prepare(hooks, c)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	expectedEfLength := 1
	actualEfLength := len(ef)
	if expectedEfLength != actualEfLength {
		t.Error("Expected a file as one file was set")
	}

	var pathToDownloadTo = fmt.Sprintf("%s/example-0", cachDir)
	if ef[0].Path != pathToDownloadTo {
		t.Error("File name is not as expected")
	}

	expectedFtcLength := 1
	actualFtcLength := len(ftc)
	if expectedFtcLength != actualFtcLength {
		t.Error("Expected a file as one file was set")
	}

	if ftc[0].Path != pathToDownloadTo {
		t.Error("File name is not as expected")
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

func Test_nonlocal_Prepare_ProgagatesPrepError(t *testing.T) {
	nl := nonlocal{}

	c := conf.Configuration{}
	unresolveableUrl := "https://example.123faketld/some_file.go"
	hooks := conf.Hooks{
		HookName: "commit-msg",
		Files: []conf.HookFile{
			conf.HookFile{
				Path: unresolveableUrl,
			},
		},
	}

	_, _, err := nl.Prepare(hooks, c)

	expectedErrorMessage := fmt.Sprintf("Failed to retrieve non-local file. Please validate that you have access to the file, and that the configured URL is correct: %s", unresolveableUrl)
	actualErrorMessage := err.Error()
	if expectedErrorMessage != actualErrorMessage {
		t.Error("Error message not as expected")
	}
}
