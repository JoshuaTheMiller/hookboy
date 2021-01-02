package generators

import (
	"fmt"
	"testing"

	"github.com/hookboy/source/hookboy"
	"github.com/hookboy/source/hookboy/conf"
	p "github.com/hookboy/source/hookboy/prep/internal"
)

func Test_Generate_ReturnsNameAsExpected(t *testing.T) {
	l := localHooksGenerator{}

	name := l.Name()

	expectedName := "Local Hook Files"
	actualName := name

	if expectedName != actualName {
		t.Error("Name was not as expected")
	}
}

func Test_Generate_ReturnsEmptyWhenAutoAddHooksIsNotSetToByFileName(t *testing.T) {
	l := localHooksGenerator{}

	ef, ftc, err := l.Generate(conf.Configuration{
		AutoAddHooks: conf.No,
	}, nil)

	expectedEfLength := 0
	actualEfLength := len(ef)
	if expectedEfLength != actualEfLength {
		t.Error("Expected ExecutableFiles array to be empty")
	}

	expectedFtcLength := 0
	actualFtcLength := len(ftc)
	if expectedFtcLength != actualFtcLength {
		t.Error("Expected FilesToCreate array to be empty")
	}

	if err != nil {
		t.Error("Expected error to be nil")
	}
}

func Test_Generate_ReturnsPropagatesError(t *testing.T) {
	l := localHooksGenerator{}

	expectedError := "Error prepping hooks by filename"

	_, _, err := l.Generate(conf.Configuration{
		AutoAddHooks: conf.ByFileName,
	}, func(string) ([]p.SimpleFile, hookboy.Error) {
		return nil, hookboy.NewError(expectedError)
	})

	if err.Error() != expectedError {
		t.Error("Expected error to be propagated")
	}
}

func Test_getHooksByFileName_PropogatesError(t *testing.T) {
	someError := hookboy.NewError("SomeError")

	errorReturningRead := func(dirname string) ([]p.SimpleFile, hookboy.Error) {
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

func Test_Generate_ReturnsFilesAsFound(t *testing.T) {
	l := localHooksGenerator{}

	localHookDirectory := "somelocaldir"

	c := conf.Configuration{
		AutoAddHooks: conf.ByFileName,
		LocalHookDir: localHookDirectory,
	}

	recognizedHook := "commit-msg"
	filesToReturn := []p.SimpleFile{
		testFile{
			name: recognizedHook,
		},
		testFile{
			// This is not a recognized git hook, and so should not be found/returned
			name: "b",
		},
	}

	fileReturningRead := func(dirname string) ([]p.SimpleFile, hookboy.Error) {
		return filesToReturn, nil
	}

	ef, ftc, err := l.Generate(c, fileReturningRead)

	if err != nil {
		t.Errorf("Expected error to be nil")
		return
	}

	expectedLengthOfFtc := 0
	actualLengthOfFtc := len(ftc)
	if expectedLengthOfFtc != actualLengthOfFtc {
		t.Errorf("Expected amount of files to create to be returned to be '%d', found '%d'", expectedLengthOfFtc, actualLengthOfFtc)
	}

	countOfFiles := len(ef)
	if countOfFiles != 1 {
		t.Errorf("Expected amount of files to be returned to be '%d', found '%d'", 1, countOfFiles)
	}

	expectedHook := recognizedHook
	actualHook := ef[0].AssociatedHook
	if actualHook != expectedHook {
		t.Errorf("Expected amount of files to be returned to be '%s', found '%s'", expectedHook, actualHook)
	}

	expectedPath := fmt.Sprintf("./%s/%s", localHookDirectory, recognizedHook)
	actualPath := ef[0].Path
	if actualPath != expectedPath {
		t.Errorf("Expected path to be '%s', found '%s'", expectedPath, actualPath)
	}
}

func Test_getHooksByFileName_ReturnsFilesAsFound(t *testing.T) {
	localHookDirectory := "somelocaldir"
	recognizedHook := "commit-msg"
	filesToReturn := []p.SimpleFile{
		testFile{
			name: recognizedHook,
		},
		testFile{
			// This is not a recognized git hook, and so should not be found/returned
			name: "b",
		},
	}

	fileReturningRead := func(dirname string) ([]p.SimpleFile, hookboy.Error) {
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
