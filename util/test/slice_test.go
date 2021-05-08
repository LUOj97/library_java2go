/*
This is a test for slice which a dynamic array in Go
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

var s []string

/*
test construct in java
*/
func TestNew(t *testing.T) {
	// declare a nil slice
	var s1 []string
	//test
	dataType := reflect.TypeOf(s1)
	fmt.Println(dataType)
	fmt.Println(s1 == nil)
}

/*
add(E e)
*/
func TestAdd(t *testing.T) {
	//add a into s(slice)
	e := "new"
	s := append(s, e)
	//test
	var stringType string
	if !reflect.DeepEqual(s[0], e) {
		t.Errorf("excepted:%v, got:%v", e, s[0])
	}
	if !reflect.DeepEqual(reflect.TypeOf(s[0]), reflect.TypeOf(stringType)) {
		t.Errorf("excepted:%v, got:%v", reflect.TypeOf(stringType), reflect.TypeOf(s[0]))
	}
}

//add(int index, E element)
func TestAddWitHIndex(t *testing.T) {
	s = make([]string, 10, 10)
	//add a into s(slice)
	s[2] = "two"
	//test
	if !reflect.DeepEqual(s[2], "two") {
		t.Errorf("excepted:%v, got:%v", "two", s[2])
	}
}

//addAll(Collection<? extends E> c)
func TestAddAll(t *testing.T) {
	sAdd := []string{"s2-1", "s2-2", "s2-3"}
	s = make([]string, 2, 10)
	for i := 0; i < len(s); i++ {
		s[i] = "s1"
	}
	//add a slice to a another slice
	sNew := append(s, sAdd...)
	//test
	if !reflect.DeepEqual(sNew[2], sAdd[0]) {
		t.Errorf("excepted:%v, got:%v", "two", s[2])
	}
	fmt.Println(sNew)
}
