package generators

import (
	"fmt"
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

func Test_PrepareStatementFile_ConstructsAsExpected(t *testing.T) {
	hookname := "somehook"
	statement := "somestatement"
	cachDir := "somecacheparentdir"
	c := conf.Configuration{
		CacheDirectory: cachDir,
	}

	ftc := prepareStatementFile(hookname, statement, c)

	if ftc.Contents != statement {
		t.Errorf("Prepared statement file contents are incorrect.")
	}

	expectedPath := fmt.Sprintf("%s/.hookboy-cache/%s-statement", cachDir, hookname)
	actualPath := ftc.Path
	if actualPath != expectedPath {
		t.Errorf("Prepared statement file's path is incorrect (e '%s', a: '%s'", expectedPath, actualPath)
	}
}
