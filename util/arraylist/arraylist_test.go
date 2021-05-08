package arraylist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.cap; actualValue != 10 {
		t.Errorf("Got %v expected %v", actualValue, 10)
	}
}

func TestNewWithInitialCapacity(t *testing.T) {
	list := NewWithInitialCapacity(2)
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	} else if actualValue := list.cap; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}

func TestAdd(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add(1)
	fmt.Println(list.size)
	fmt.Println(list.cap)
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

func TestAddWithIndex(t *testing.T) {
	list := New()
	list.Add("0")
	list.Add("2")
	list.Add("3")
	x := "1new"
	list.AddWithIndex(1, x)
	list.AddWithIndex(4, "5new")
	list.AddWithIndex(7, "new8")

	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 5 {
		t.Errorf("Got %v expected %v", actualValue, 5)
	}
	if actualValue, ok := list.Get(1); actualValue != x || !ok {
		t.Errorf("Got %v expected %v", actualValue, x)
	}

}

func TestAddAll(t *testing.T) {
	list := New()
	list.Add(1)
	list2 := New()
	list2.Add(2)
	list2.Add(3)
	list.AddAll(list2)
	list3 := New()
	list3.Add("hello")
	list.AddAll(list3)
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, "33")
	}
	if actualValue := list.cap; actualValue != 10 {
		t.Errorf("Got %v expected %v", actualValue, "10")
	}

}

func TestAddAllWithIndex(t *testing.T) {
	list1 := New()
	list1.Add("A-0")
	list1.Add("A-1")
	list1.Add("A-2")
	list2 := New()
	list2.Add("B-0")
	list2.Add("B-1")
	list1.AddAllWithIndex(1, list2)
	listError := New()
	listError.Add(1)
	list1.AddAllWithIndex(0, listError)
	fmt.Println(list1)
	if actualValue := list1.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list1.Size(); actualValue != 5 {
		t.Errorf("Got %v expected %v", actualValue, 5)
	}
	if actualValue, ok := list1.Get(1); actualValue != "B-0" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "B-0")
	}

}

func TestClear(t *testing.T) {
	list := New()
	list.Add("e")
	list.Add("b")
	list.Clear()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestContains(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	if actualValue := list.Contains("a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("b"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	list.Clear()
	if actualValue := list.Contains("d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	type student struct {
		name string
		age  int
	}
	s1 := student{"bob", 11}
	s2 := student{"alice", 11}
	s2a := student{"alice", 11}
	s3 := student{"error", 11}
	list2 := New()
	list2.Add(s1)
	list2.Add(s2)
	fmt.Println(list2)
	if actualValue := list2.Contains(s1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list2.Contains(s2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list2.Contains(s2a); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list2.Contains(s3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	//list.Remove(0)
	//if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
	//	t.Errorf("Got %v expected %v", actualValue, "b")
	//}
}

func TestIndexOf(t *testing.T) {
	list := New()
	expectedIndex := -1
	if index := list.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
	list.Add("a")
	list.Add("b")

	expectedIndex = 0
	if index := list.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 1
	if index := list.IndexOf("b"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

}

func TestIsEmpty(t *testing.T) {
	list := New()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestLastIndexOf(t *testing.T) {
	list := New()
	expectedIndex := -1
	if index := list.LastIndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
	list.Add("a")
	list.Add("b")
	list.Add("a")
	list.Add("b")
	list.Add("a")
	expectedIndex = 4
	if index := list.LastIndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 3
	if index := list.LastIndexOf("b"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

}

func TestRemoveWithIndex(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.RemoveWithIndex(1)

	if actualValue, _ := list.Get(1); actualValue != "c" {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	list.RemoveWithIndex(1)
	list.RemoveWithIndex(0)
	list.RemoveWithIndex(0) // no effect
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestRemoveWithObeject(t *testing.T) {
	type student struct {
		name string
		age  int
	}
	s1 := student{"bob", 11}
	s2 := student{"alice", 12}
	s3 := student{"error", 15}
	s2a := student{"alice", 12}
	list := New()
	list.Add(s1)
	list.Add(s2)
	list.Add(s3)
	list.Add(s2a)
	list.RemoveWithObeject(s2a)
	fmt.Println(list)
	if actualValue, _ := list.Get(1); actualValue != s3 {
		t.Errorf("Got %v expected %v", actualValue, s3)
	}
	if actualValue, _ := list.Get(2); actualValue != s2a {
		t.Errorf("Got %v expected %v", actualValue, s2a)
	}
	list.RemoveWithObeject(s2a)
	list.RemoveWithObeject(s1)
	list.RemoveWithObeject(s3)
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSet(t *testing.T) {
	list := New()
	list.Set(0, "a")
	list.Set(1, "b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.Set(2, "c") // append
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Set(4, "d")  // ignore
	list.Set(1, "bb") // update
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", list.ToArray()...), "abbc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSize(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.RemoveWithIndex(1)
	if actualValue := list.Size(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

}

func TestSubList(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	if actualValue := list.Size(); actualValue != 5 {
		t.Errorf("Got %v expected %v", actualValue, 5)
	}
	sublist := list.SubList(1, 3)
	if actualValue := sublist.Size(); actualValue != 3-1 {
		t.Errorf("Got %v expected %v", actualValue, 3-1)
	}

	if actualValue, ok := sublist.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3-1)
	}
	if actualValue, ok := sublist.Get(1); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3-1)
	}

}

func TestToArray(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", list.ToArray()...), "abc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestTrimToSize(t *testing.T) {
	list := NewWithInitialCapacity(10)
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.TrimToSize()
	if actualValue := cap(list.elements); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.RemoveWithIndex(0)
	list.TrimToSize()
	if actualValue := cap(list.elements); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}
