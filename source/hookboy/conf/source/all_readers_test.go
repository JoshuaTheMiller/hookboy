package source

import (
	"testing"

	"github.com/hookboy/source/hookboy/conf/deserialization"
)

var packageJSONContents = `{
	"hookboy":{
		"localHookDir":"./someRandomHooksDirForTestingPackageJSON"
	}
}`
var packageJSONOption = fileSystemObjectOptions{
	Name:         "package.json",
	FileContents: packageJSONContents,
}

var yamlContents = `---

localHookDir: ./someRandomHooksDirForTestingYAML
`
var packageYamlOption = fileSystemObjectOptions{
	Name:         "hookboy.yml",
	FileContents: yamlContents,
}

var folderOptions = fileSystemObjectOptions{
	Name:     ".hookboy",
	IsFolder: true,
}

type sourceTestOptions struct {
	FileContents              string
	ConfigValueToCheckFor     string
	FileSystemCreationOptions []fileSystemObjectOptions
	ReaderToTest              configurationReader
}

var testsToRun = map[string]sourceTestOptions{
	"source_package_json": sourceTestOptions{
		FileContents:              packageJSONContents,
		ConfigValueToCheckFor:     "./someRandomHooksDirForTestingPackageJSON",
		FileSystemCreationOptions: []fileSystemObjectOptions{packageJSONOption},
		ReaderToTest:              packageJSONFileReader{},
	},
	"source_local_file": sourceTestOptions{
		FileContents:              yamlContents,
		ConfigValueToCheckFor:     "./someRandomHooksDirForTestingYAML",
		FileSystemCreationOptions: []fileSystemObjectOptions{packageYamlOption},
		ReaderToTest: localFileReader{
			Path:         packageYamlOption.Name,
			Deserializer: deserialization.YamlDeserializer{},
		},
	},
	"source_local_folder": sourceTestOptions{
		FileContents:              "",
		ConfigValueToCheckFor:     "",
		FileSystemCreationOptions: []fileSystemObjectOptions{folderOptions},
		ReaderToTest: localFolderReader{
			Path: folderOptions.Name,
		},
	},
}

func TestAllCanRead(t *testing.T) {
	for readerUnderTestName, opts := range testsToRun {
		testCanRead(t, readerUnderTestName, opts)
	}
}

func testCanRead(t *testing.T, readerUnderTestName string, options sourceTestOptions) {
	var fileSystemObjectOptions = options.FileSystemCreationOptions
	creatFileSystemObjectForTest(fileSystemObjectOptions...)
	defer deleteFileSystemObjectForTest(fileSystemObjectOptions...)

	var reader = options.ReaderToTest

	var objectToRead = fileSystemObjectOptions[0].Name

	var readerCanRead = reader.CanRead()

	if !readerCanRead {
		t.Errorf("%s: Expected reader to be able to find %s, did not", readerUnderTestName, objectToRead)
		return
	}
}

func TestCanReadDoesNotFind(t *testing.T) {
	for readerUnderTestName, opts := range testsToRun {
		testCanNotRead(t, readerUnderTestName, opts)
	}
}

func testCanNotRead(t *testing.T, readerUnderTestName string, options sourceTestOptions) {
	var fileSystemObjectOptions = options.FileSystemCreationOptions

	var reader = options.ReaderToTest

	var objectToRead = fileSystemObjectOptions[0].Name

	var readerCanRead = reader.CanRead()

	if readerCanRead {
		t.Errorf("%s: Expected reader to not find %s, did", readerUnderTestName, objectToRead)
		return
	}
}

func TestReadReturnsExpectedConfiguration(t *testing.T) {
	for readerUnderTestName, opts := range testsToRun {
		testReadReturnsExpectedConfiguration(t, readerUnderTestName, opts)
	}
}

func testReadReturnsExpectedConfiguration(t *testing.T, readerUnderTestName string, options sourceTestOptions) {
	var fileSystemObjectOptions = options.FileSystemCreationOptions
	creatFileSystemObjectForTest(fileSystemObjectOptions...)
	defer deleteFileSystemObjectForTest(fileSystemObjectOptions...)

	var reader = options.ReaderToTest

	var configuration, err = reader.Read()

	if err != nil {
		t.Errorf("%s: Reader returned error when success was expected: %s", readerUnderTestName, err)
		return
	}

	// .someRandomHooksDirForTesting comes from the hardcoded json value above
	if configuration.LocalHookDir != options.ConfigValueToCheckFor {
		t.Errorf("%s: Reader returned an unexpected configuration value. This could mean there is an issue with the deserializers.", readerUnderTestName)
		return
	}
}
