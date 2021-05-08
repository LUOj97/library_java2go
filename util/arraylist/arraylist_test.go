package arraylist

import (
	"testing"
)

func TestListNew(t *testing.T) {
	list := New()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestListNewWithInitialCapacity(t *testing.T) {
	list := NewWithInitialCapacity(2)
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	} else if actualValue := list.cap; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}

func TestListAdd(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue := list.cap; actualValue != 10 {
		t.Errorf("Got %v expected %v", actualValue, "10")
	}
}

func TestListAddAll(t *testing.T) {
	list := New()
	list.Add("11")
	list2 := New()
	list2.Add("22")
	list2.Add("33")
	list.AddAll(list2)
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "33" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "33")
	}
	if actualValue := list.cap; actualValue != 10 {
		t.Errorf("Got %v expected %v", actualValue, "10")
	}
}
