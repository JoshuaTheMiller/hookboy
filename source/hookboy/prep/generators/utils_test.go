package generators

import (
	"testing"
)

func Test_itemExists_ReturnsTrueWhenItemDoesIndeedExist(t *testing.T) {
	var someItem = "what"
	var someArray = [...]string{someItem, "what2"}

	// itemExists is not very clear as far as what param goes where
	// #CopyPasteFails
	var itemExists = itemExists(someArray, someItem)

	if !itemExists {
		t.Error("Item existed, but not found")
	}
}

func Test_itemExists_PanicsWhenNotArray(t *testing.T) {
	var someItem = "what"
	var someArray = "thisisastring"

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Expected panic, did not appear to panic")
		}

		var exectedMessage = "Invalid data-type"
		if err != exectedMessage {
			t.Errorf("e: '%s', a: '%s'", exectedMessage, err)
		}
	}()

	itemExists(someArray, someItem)
}
