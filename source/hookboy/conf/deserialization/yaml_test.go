package deserialization

import (
	"testing"

	"github.com/hookboy/source/hookboy/conf"
)

var testData1 = `
localHookDir: ./somethingElse
autoAddHooks: No
hooks: []
`

func TestAutoAddHooksSetToNoneProperly(t *testing.T) {
	var configuration, _ = YamlDeserializer{}.Deserialize([]byte(testData1))

	if configuration.AutoAddHooks != conf.No {
		t.Errorf("Expected AutoAddHooks to be No")
	}

	var expectedLocalHooksDir = "./somethingElse"
	if configuration.LocalHookDir != expectedLocalHooksDir {
		t.Errorf("LocalHookDir, expected %s, got %s", expectedLocalHooksDir, configuration.LocalHookDir)
	}
}

var testData2 = `
localHookDir: []
autoAddHooks: What
hooks: []
`

func TestInvalidConfigDataTreatedAsError(t *testing.T) {
	var _, deserializationError = YamlDeserializer{}.Deserialize([]byte(testData2))

	var expectedErrorMessage = "failed to parse configuration file"
	var actualErrorMessage = deserializationError.Error()
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message to be '%s', got '%s'", expectedErrorMessage, actualErrorMessage)
	}
}
