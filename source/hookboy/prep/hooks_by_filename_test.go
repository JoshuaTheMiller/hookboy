package prep

import (
	"errors"
	"fmt"
	"testing"
)

func Test_getHooksByFileName_PropogatesError(t *testing.T) {
	someError := errors.New("SomeError")

	errorReturningRead := func(dirname string) ([]simpleFile, error) {
		return nil, someError
	}

	files, err := getHooksByFileName("na", errorReturningRead)

	if err == nil {
		t.Errorf("Expected error to be returned")
		return
	}

	if err != someError {
		t.Errorf("e: '%s', a: '%s", someError, err)
	}

	if files != nil {
		t.Errorf("Expected files to be nil")
	}
}

type testFile struct {
	name string
}

func (t testFile) Name() string {
	return t.name
}

func Test_getHooksByFileName_ReturnsFilesAsFound(t *testing.T) {
	localHookDirectory := "somelocaldir"
	recognizedHook := "commit-msg"
	filesToReturn := []simpleFile{
		testFile{
			name: recognizedHook,
		},
		testFile{
			// This is not a recognized git hook, and so should not be found/returned
			name: "b",
		},
	}

	fileReturningRead := func(dirname string) ([]simpleFile, error) {
		return filesToReturn, nil
	}

	files, err := getHooksByFileName(localHookDirectory, fileReturningRead)

	if err != nil {
		t.Errorf("Expected error to be nil")
		return
	}

	countOfFiles := len(files)
	if countOfFiles != 1 {
		t.Errorf("Expected amount of files to be returned to be '%d', found '%d'", 1, countOfFiles)
	}

	expectedHook := recognizedHook
	actualHook := files[0].AssociatedHook
	if actualHook != expectedHook {
		t.Errorf("Expected amount of files to be returned to be '%s', found '%s'", expectedHook, actualHook)
	}

	expectedPath := fmt.Sprintf("./%s/%s", localHookDirectory, recognizedHook)
	actualPath := files[0].Path
	if actualPath != expectedPath {
		t.Errorf("Expected path to be '%s', found '%s'", expectedPath, actualPath)
	}
}
