package runner

import (
	"os"
	"testing"
)

func setup() {
	os.MkdirAll(".git/hooks", os.ModeDir)
	os.MkdirAll("hooks", os.ModeDir)
}

func tearDown() {
	os.RemoveAll(".git")
	os.RemoveAll("hooks")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}
