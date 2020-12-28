package prep

import "reflect"

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

type mapItem struct {
	Key   interface{}
	Value interface{}
}

type mappable interface {
	ToMapItem() mapItem
}

func toMapOfGroups(arr ...mappable) map[interface{}][]interface{} {
	var m = make(map[interface{}][]interface{})

	for _, item := range arr {
		mapped := item.ToMapItem()

		var key = mapped.Key
		var value = mapped.Value

		existingItems, exists := m[key]

		if exists {
			existingItems = append(existingItems, value)
			m[key] = existingItems
		} else {
			m[key] = []interface{}{value}
		}
	}

	return m
}

type groupable interface {
	SelectKey() interface{}
	SelectValue() interface{}
}

func groupBy(a []groupable) map[interface{}][]interface{} {
	var m = make(map[interface{}][]interface{})

	for _, item := range a {
		var key = item.SelectKey()
		var value = item.SelectValue()

		existingItems, exists := m[key]

		if exists {
			existingItems = append(existingItems, value)
			m[key] = existingItems
		} else {
			m[key] = []interface{}{value}
		}
	}

	return m
}

type grouping interface {
	Get(k interface{}) []interface{}
	Count() int
	Iter(fn func(key interface{}, values interface{}))
}

type actualGrouping struct {
	internals map[interface{}][]interface{}
}

func (g actualGrouping) Get(k interface{}) []interface{} {
	return g.internals[k]
}

func (g actualGrouping) Count() int {
	return len(g.internals)
}

func (g actualGrouping) Iter(fn func(key interface{}, values interface{})) {
	for key, values := range g.internals {
		fn(key, values)
	}
}

func groupByKey(a interface{}, keySelector func(a interface{}) (b interface{}), valueSelector func(a interface{}) (c interface{})) grouping {
	var m = make(map[interface{}][]interface{})

	for _, item := range a.([]interface{}) {
		var key = keySelector(item)
		var value = valueSelector(item)

		existingItems, exists := m[key]

		if exists {
			existingItems = append(existingItems, value)
			m[key] = existingItems
		} else {
			m[key] = []interface{}{value}
		}
	}

	return actualGrouping{
		internals: m,
	}
}
