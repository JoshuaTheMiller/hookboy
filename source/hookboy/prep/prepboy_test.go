package prep

import (
	"fmt"
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

func Test_PrepareHookfileInfo_PreparesNoFilesForGenerationWhenThereAreNoneToGenerate(t *testing.T) {
	// TODO: control ReadDir to mitigate possible found hooks
	p := prepboy{}
	c := conf.Configuration{}
	ftc, err := p.PrepareHookfileInfo(c)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	expectedAmountOfFilesToGenerate := 0
	actualAmountOfFilesToGenerate := len(ftc)

	if actualAmountOfFilesToGenerate != expectedAmountOfFilesToGenerate {
		t.Errorf("Expected the amount of files to generate to be 0, given that none were found or set via configuration")
		return
	}
}

func Test_PrepareHookfileInfo_Prepares2FilesAsExpected(t *testing.T) {
	p := prepboy{}
	c := conf.Configuration{
		AutoAddHooks: conf.No,
		Hooks: []conf.Hooks{
			conf.Hooks{
				// Will prepare the hookfile itself
				HookName: "commit-msg",
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
	ftc, err := p.PrepareHookfileInfo(c)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	expectedAmountOfFilesToGenerate := 2
	actualAmountOfFilesToGenerate := len(ftc)

	if actualAmountOfFilesToGenerate != expectedAmountOfFilesToGenerate {
		t.Errorf("Expected the amount of files to generate to be 2, one for the hookfile, one for the included statement")
		return
	}
}

func Test_PrepareHookfileInfo_PropgatesErrorWhenTryingToReadNonExistantFolder(t *testing.T) {
	p := prepboy{}
	c := conf.Configuration{
		AutoAddHooks: conf.ByFileName,
	}

	_, err := p.PrepareHookfileInfo(c)

	if err == nil {
		t.Error("Expected error to be present")
		return
	}

	prepboyError, ok := err.(prepboyError)
	if ok != true {
		t.Error("Expected error to be prepboyError")
		return
	}

	expectedDescription := "Error prepping hooks by filename"
	if prepboyError.Error() != expectedDescription {
		t.Error("Expected description does not match actual")
	}
}

func Test_PrepareHookfileInfo_Prepares1FileWhenAutoAdding1File(t *testing.T) {
	c := conf.Configuration{
		AutoAddHooks: conf.ByFileName,
	}
	hookfileForTest := simpleFileForTest{
		name: "commit-msg",
	}
	p := prepboy{
		ReadDir: func(dirname string) ([]simpleFile, error) {
			return []simpleFile{
				hookfileForTest,
			}, nil
		},
		Instantiated: true,
	}

	ftc, err := p.PrepareHookfileInfo(c)

	if err != nil {
		t.Error("Expected error to be nil")
		return
	}

	expectedAmountOfFilesToGenerate := 1
	actualAmountOfFilesToGenerate := len(ftc)

	if actualAmountOfFilesToGenerate != expectedAmountOfFilesToGenerate {
		t.Errorf("Expected the amount of files to generate to be 1: one for the hookfile")
		return
	}
}

func Test_Instantiate_ReturnsInstantiatedPrepboy(t *testing.T) {
	prepboy := prepboy{}

	prepboy.instantiate()

	if prepboy.Instantiated != true {
		t.Error("Expected prepboy to be instantiated")
	}

	if prepboy.ReadDir == nil {
		t.Error("Expected ReadDir to be set")
	}
}

func Test_readDir_ReturnsExpectedInformation(t *testing.T) {
	dir := "testbin"
	filename := "somefile"
	filepath := fmt.Sprintf("./%s/%s", dir, filename)
	contents := "what"
	fso := []fileSystemObjectOptions{
		fileSystemObjectOptions{
			Name:     dir,
			IsFolder: true,
		},
		fileSystemObjectOptions{
			Name:         filepath,
			FileContents: contents,
		},
	}
	creatFileSystemObjectForTest(fso...)
	defer deleteFileSystemObjectForTest(fso...)

	actualObject, err := readDir(dir)

	if err != nil {
		t.Error("Found error when none was expected")
		return
	}

	expectedCountOfFiles := 1
	actualCountOfFiles := len(actualObject)
	if expectedCountOfFiles != actualCountOfFiles {
		t.Errorf("Expected %d file, found %d", expectedCountOfFiles, actualCountOfFiles)
		return
	}

	expectedFileName := filename
	actualFileName := actualObject[0].Name()
	if expectedFileName != actualFileName {
		t.Errorf("Incorrect  name. e: '%s', a: '%s'", expectedFileName, actualFileName)
	}
}
