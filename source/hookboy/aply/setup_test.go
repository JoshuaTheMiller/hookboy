package aply

import (
	"os"
	"testing"
)

const (
	FileMode = os.ModePerm
)

func setup() {
	os.MkdirAll(".git/hooks", FileMode)
	os.MkdirAll("hooks", FileMode)
}

func tearDown() {
	os.RemoveAll(".git")
	os.RemoveAll("hooks")
	os.RemoveAll(".grapple-cache")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}
