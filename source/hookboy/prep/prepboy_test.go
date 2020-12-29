package prep

import (
	"fmt"
	"testing"
)

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
