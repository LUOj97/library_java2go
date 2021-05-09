package hashmap

import (
	"fmt"
	"testing"
)

func TestMapPut(t *testing.T) {
	m := New()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put("1", "a") //overwrite
	fmt.Println(m)
}

func TestHashMap_Clear(t *testing.T) {
	m := New()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	m.Clear()
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

}
func TestHashMap_ContainsKey(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	m.Put(4, "c")
	m.Put(5, "d")
	m.Put(6, "b")

	if actualValue := m.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}
	tests2 := [][]interface{}{
		{1, "a", true},
		{1, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, 5, true},
		{10, nil, false},
		{"2", nil, false},
		{"1", "e", false},
	}

	for _, test := range tests2 {
		ok := m.ContainsKey(test[0])
		if ok != test[2] {
			t.Errorf("Got %v expected %v", ok, test[2])
		}
	}

}
func TestHashMap_ContainsValue(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	m.Put(4, "c")
	m.Put(5, "d")
	m.Put(6, "b")

	if actualValue := m.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}
	tests2 := [][]interface{}{
		{1, "a", false},
		{1, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, 5, false},
		{10, nil, false},
		{"2", nil, false},
		{"1", "e", true},
	}

	for _, test := range tests2 {
		ok := m.ContainsValue(test[1])
		if ok != test[2] {
			t.Errorf("Got %v expected %v", ok, test[2])
		}
	}

}
func TestHashMap_Get(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	m.Put(4, "c")
	m.Put(5, "d")
	m.Put(6, "b")

	if actualValue := m.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}
	tests2 := [][]interface{}{
		{1, "e"},
		{2, "f"},
		{"1", nil},
		{3, "g"},
		{4, "c"},
	}

	for _, test := range tests2 {
		ok := m.Get(test[0])
		if ok != test[1] {
			t.Errorf("Got %v expected %v", ok, test[1])
		}
	}

}
func TestHashMap_GetOrDefault(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	m.Put(4, "c")
	m.Put(5, "d")
	m.Put(6, "b")

	if actualValue := m.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}
	tests2 := [][]interface{}{
		{1, "e"},
		{2, "f"},
		{3, "g"},
		{4, "c"},
		{"1", "default"},
		{"aaaa", "default"},
	}

	for _, test := range tests2 {
		ok := m.GetOrDefault(test[0], "default")
		if ok != test[1] {
			t.Errorf("Got %v expected %v", ok, test[1])
		}
	}

}
func TestHashMap_IsEmpty(t *testing.T) {
	m := New()
	if actualValue := m.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	if actualValue := m.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

}
func TestHashMap_Put(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	m.Put(4, "c")
	m.Put(5, "d")
	m.Put(6, "b")
	if actualValue := m.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}
	tests2 := [][]interface{}{
		{1, "e"},
		{2, "f"},
		{3, "g"},
		{4, "c"},
		{"4", nil},
	}
	for _, test := range tests2 {
		value := m.Get(test[0])
		if value != test[1] {
			t.Errorf("Got %v expected %v", value, test[1])
		}
	}

}
func TestHashMap_PutIfAbsent(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	if actualValue := m.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	tests2 := [][]interface{}{
		{m.PutIfAbsent(1, "e"), "e"},
		{m.PutIfAbsent(1, 1), "e"},
		{m.PutIfAbsent(2, "f"), "f"},
		{m.PutIfAbsent(4, "four"), nil},
		{m.PutIfAbsent(5, "four"), nil},
		{m.PutIfAbsent(6, "sixe"), nil},
	}
	for _, test := range tests2 {
		value := test[0]
		if value != test[1] {
			t.Errorf("Got %v expected %v", value, test[1])
		}
	}
}

func TestHashMap_Remove(t *testing.T) {
	m := New()
	m.Put("1", "e")
	m.Put("a", "f")
	m.Put("c", "g")
	if actualValue := m.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	tests2 := [][]interface{}{
		{m.Remove("1"), "e"},
		{m.Remove("1"), nil},
		{m.Remove("a"), "f"},
		{m.Remove("a"), nil},
		{m.Remove("c"), "g"},
		{m.Remove("aaaaa"), nil},
	}
	for _, test := range tests2 {
		value := test[0]
		if value != test[1] {
			t.Errorf("Got %v expected %v", value, test[1])
		}
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}
func TestHashMap_RemoveWithKeyValue(t *testing.T) {
	m := New()
	m.Put("1", "e")
	m.Put("a", "f")
	m.Put("c", "g")
	if actualValue := m.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	tests2 := [][]interface{}{
		{m.RemoveWithKeyValue("1", "aaaaaa"), false},
		{m.RemoveWithKeyValue("1", "e"), true},
		{m.RemoveWithKeyValue("a", "a"), false},
		{m.RemoveWithKeyValue("a", "f"), true},
		{m.RemoveWithKeyValue("l", "g"), false},
		{m.RemoveWithKeyValue("c", "g"), true},
	}
	for _, test := range tests2 {
		value := test[0]
		if value != test[1] {
			t.Errorf("Got %v expected %v", value, test[1])
		}
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestHashMap_Replace(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2.0, "f")
	m.Put(3.0, "g")
	if actualValue := m.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	tests2 := [][]interface{}{
		{m.Replace(1, "aaaaaa"), "e"},
		{m.Get(1), "aaaaaa"},
		{m.Replace("1", "e"), nil},
		{m.Replace(2.0, "a"), "f"},
		{m.Get(2.0), "a"},
	}
	for _, test := range tests2 {
		value := test[0]
		if value != test[1] {
			t.Errorf("Got %v expected %v", value, test[1])
		}
	}

}
func TestHashMap_ReplaceWithValue(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2.0, "f")
	m.Put(3.0, "g")
	if actualValue := m.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	tests2 := [][]interface{}{
		{m.ReplaceWithValue(1, "e", "1new"), true},
		{m.ReplaceWithValue(1, "error", "1new"), false},
		{m.Get(1), "1new"},
		{m.ReplaceWithValue(3.0, "error", "3.0new"), false},
		{m.ReplaceWithValue(3.0, "g", "3.0new"), true},
		{m.Get(3.0), "3.0new"},
	}
	for _, test := range tests2 {
		value := test[0]
		if value != test[1] {
			t.Errorf("Got %v expected %v", value, test[1])
		}
	}

}

func TestHashMap_Size(t *testing.T) {
	m := New()
	m.Put(1, "e")
	m.Put(2, "f")
	m.Put(3, "g")
	if actualValue := m.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

}
