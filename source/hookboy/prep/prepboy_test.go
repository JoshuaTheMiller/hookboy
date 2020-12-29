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

func Test_Instantiate_ReturnsInstantiatedPrepboy(t *testing.T) {
	c := conf.Configuration{
		CacheDirectory: "somedir",
	}
	prepboy := Instantiate(c).(prepboy)

	if prepboy.Instantiated != true {
		t.Error("Expected prepboy to be instantiated")
	}

	if prepboy.C.CacheDirectory != c.CacheDirectory {
		t.Error("Expected configuration to be the same")
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
