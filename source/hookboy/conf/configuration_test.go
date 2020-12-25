package conf

import "testing"

func TestDefaultsGetSet(t *testing.T) {
	var c = Configuration{}

	c.SetDefaults()

	if c.LocalHookDir != "./hooks" {
		t.Errorf("Expected CacheDirectory to be './hooks', got '%s'", c.LocalHookDir)
	}

	if c.AutoAddHooks != ByFileName {
		t.Errorf("Expected AutoAddHooks to be 'ByFileName', got '%s'", c.AutoAddHooks)
	}

	if c.CacheDirectory != ".git/hooks/" {
		t.Errorf("Expected CacheDirectory to be '.git/hooks/', got '%s'", c.CacheDirectory)
	}
}

func TestGetCacheDirectory(t *testing.T) {
	var cacheParentFolder = "someDir"
	var c = Configuration{
		CacheDirectory: cacheParentFolder,
	}

	var actual = c.GetCacheDirectory()
	var expectedValue = "someDir/.hookboy-cache"
	if actual != expectedValue {
		t.Errorf("Expected value to be '%s', got '%s'", expectedValue, actual)
	}
}
