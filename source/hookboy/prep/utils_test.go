package prep

import "testing"

type testObject struct {
	Type string
	Name string
}

func (t testObject) ToMapItem() mapItem {
	return mapItem{
		Key:   t.Type,
		Value: t,
	}
}

func TestToMapOfGroups(t *testing.T) {
	var someObjects = []mappable{
		testObject{
			Type: "A",
			Name: "1",
		},
		testObject{
			Type: "A",
			Name: "2",
		},
		testObject{
			Type: "B",
			Name: "1",
		},
	}

	var mapOfSomeObjects = toMapOfGroups(someObjects...)

	if len(mapOfSomeObjects) != 2 {
		t.Errorf("Expected map to contain 2 groups")
	}

	var aLength = len(mapOfSomeObjects["A"])
	if aLength != 2 {
		t.Errorf("Expected group of 'A' to contain 2 items, contains %d", aLength)
	}

	if len(mapOfSomeObjects["B"]) != 1 {
		t.Errorf("Expected group of 'B' to contain 1 item")
	}
}

func (t testObject) SelectKey() interface{} {
	return t.Type
}

func (t testObject) SelectValue() interface{} {
	return t
}

func TestGroupBy(t *testing.T) {
	var someObjects = []groupable{
		testObject{
			Type: "A",
			Name: "1",
		},
		testObject{
			Type: "A",
			Name: "2",
		},
		testObject{
			Type: "B",
			Name: "1",
		},
	}

	var group = groupBy(someObjects)

	if len(group) != 2 {
		t.Errorf("Expected map to contain 2 groups")
	}

	var aLength = len(group["A"])
	if aLength != 2 {
		t.Errorf("Expected group of 'A' to contain 2 items, contains %d", aLength)
	}

	if len(group["B"]) != 1 {
		t.Errorf("Expected group of 'B' to contain 1 item")
	}
}

// This would be really awesome with generics
func TestGroupByKey(t *testing.T) {
	var someObjects = []interface{}{
		testObject{
			Type: "A",
			Name: "1",
		},
		testObject{
			Type: "A",
			Name: "2",
		},
		testObject{
			Type: "B",
			Name: "1",
		},
	}

	var keySelect = func(o interface{}) interface{} { return o.(testObject).Type }
	var valueSelect = func(o interface{}) interface{} { return o }

	var grouping = groupByKey(someObjects, keySelect, valueSelect)

	if grouping.Count() != 2 {
		t.Errorf("Expected map to contain 2 groups")
	}

	var aLength = len(grouping.Get("A"))
	if aLength != 2 {
		t.Errorf("Expected group of 'A' to contain 2 items, contains %d", aLength)
	}

	if len(grouping.Get("B")) != 1 {
		t.Errorf("Expected group of 'B' to contain 1 item")
	}
}
