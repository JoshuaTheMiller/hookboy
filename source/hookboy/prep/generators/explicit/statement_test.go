package explicit

import (
	"fmt"
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

func Test_statement_Name_IsSetAsExpected(t *testing.T) {
	s := statement{}

	name := s.Name()

	expectedName := "statement"
	actualName := name

	if expectedName != actualName {
		t.Error("Name mismatch")
	}
}

func Test_statement_ReturnEmptyArraysWhenNoStatementIsPresent(t *testing.T) {
	s := statement{}

	hookname := "somehook"
	hooks := conf.Hooks{
		HookName: hookname,
	}
	c := conf.Configuration{}

	ef, ftc, err := s.Prepare(hooks, c)

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

func Test_PrepareStatementFile_ConstructsAsExpected(t *testing.T) {
	s := statement{}

	hookname := "somehook"
	statement := "somestatement"
	cachDir := "somecacheparentdir"
	hooks := conf.Hooks{
		HookName:  hookname,
		Statement: statement,
	}
	c := conf.Configuration{
		CacheDirectory: cachDir,
	}

	ef, ftc, _ := s.Prepare(hooks, c)

	expectedEfLength := 1
	actualEfLength := len(ef)
	if expectedEfLength != actualEfLength {
		t.Errorf("Expected only 1 EF, as a statement will be stored in only one file")
	}

	expectedFtcLength := 1
	actualFtcLength := len(ftc)
	if expectedFtcLength != actualFtcLength {
		t.Errorf("Expected only 1 FTC, as a statement will be stored in only one file")
	}

	if ftc[0].Contents != statement {
		t.Errorf("Prepared statement file contents are incorrect.")
	}

	expectedPath := fmt.Sprintf("%s/%s-statement", cachDir, hookname)
	actualPath := ftc[0].Path
	if actualPath != expectedPath {
		t.Errorf("Prepared statement file's path is incorrect (e '%s', a: '%s'", expectedPath, actualPath)
	}
}
