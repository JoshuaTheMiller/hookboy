package installer

import (
	"testing"
)

func TestGetHomeDirDoesNotFail(t *testing.T) {
	// Cannot actually guarantee the exact path of the home dir,
	// without depending on external packages, which is somewhat
	// annoying so we're going to test for general success.
	// Open to suggestions though!

	_, err := GetHomeDir()

	if err != nil {
		t.Errorf("Unable to retrieve HomeDir: %s", err)
	}
}
