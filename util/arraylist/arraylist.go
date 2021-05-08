// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraylist implements the array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
//Reference:https://github.com/emirpasic/gods/blob/master/lists/arraylist/arraylist.go
package arraylist

import (
	"errors"
	"fmt"
	"github.com/emirpasic/gods/utils"
	"reflect"
	"strings"
)

// List holds the elements in a slice
type ArrayList struct {
	elements []interface{}
	size     int
	cap      int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// ArrayList()
func New() *ArrayList {
	list := &ArrayList{
		elements: make([]interface{}, 0, 10),
		size:     0,
		cap:      10,
	}
	return list
}

//lacking
//ArrayList(Collection<? extends E> c)

// ArrayList(int initialCapacity)
func NewWithInitialCapacity(capacity int) *ArrayList {
	list := &ArrayList{
		elements: make([]interface{}, 0, capacity),
		size:     0,
		cap:      capacity,
	}
	return list
}

// add(E e)
func (list *ArrayList) Add(values interface{}) error {
	errorMsg := list.checkDataType(values)
	if errorMsg == nil {
		list.elements = append(list.elements, values)
		list.size = len(list.elements)
		list.cap = cap(list.elements)
		return nil
	} else {
		fmt.Println(errorMsg)
		return errors.New("data type error")
	}
}

func (list *ArrayList) checkDataType(values interface{}) error {
	if len(list.elements) == 0 || reflect.TypeOf(list.elements[list.size-1]) == reflect.TypeOf(values) {
		return nil
	} else {
		return errors.New("data type error")
	}
}

//add(int index, E element)
func (list *ArrayList) AddWithIndex(index int, values interface{}) error {
	errorMsg := list.checkDataType(values)
	if errorMsg == nil {
		if !list.withinRange(index) {
			if index == list.size {
				list.Add(values)
				return nil
			} else {
				return errors.New("index error")
			}

		} else {
			list.growBy(1, 1)
			copy(list.elements[index+1:], list.elements[index:list.size-1])
			list.elements[index] = values
			return nil
		}
	} else {
		fmt.Println(errorMsg)
		return errors.New("data type error")
	}
}

//addAll(int index, Collection<? extends E> c)
func (list *ArrayList) AddAll(e *ArrayList) {
	if list.checkDataType(e.elements[0]) == nil {
		list.elements = append(list.elements, e.elements...)
		list.size = len(list.elements)
		list.cap = cap(list.elements)
	}
}

//addAll(int index, Collection<? extends E> c)
func (list *ArrayList) AddAllWithIndex(index int, e *ArrayList) error {
	errorMsg := list.checkDataType(e.elements[0])
	if errorMsg == nil {
		if !list.withinRange(index) {
			if index == list.size {
				list.AddAll(e)
				return nil
			} else {
				return errors.New("index error")
			}
		} else {
			l := e.size
			c := e.cap
			list.growBy(l, c)
			copy(list.elements[index+l:], list.elements[index:list.size-l])
			copy(list.elements[index:], e.elements)
			return nil
		}
	} else {
		fmt.Println(errorMsg)
		return errors.New("data type error")
	}
}

//clear()
func (list *ArrayList) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

//contains(Object o)
func (list *ArrayList) Contains(values interface{}) bool {
	var flag bool = false
	if list.IndexOf(values) != -1 {
		flag = true
	}
	return flag
}

//ensureCapacity(int minCapacity)
//实现

//forEach(Consumer<? super E> action)
//实现

//get(int index)
func (list *ArrayList) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

//indexOf(Object o)
func (list *ArrayList) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if reflect.DeepEqual(element, value) {
			return index
		}
	}
	return -1
}

//isEmpty()
func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

//iterator()
//实现

//lastIndexOf(Object o)
func (list *ArrayList) LastIndexOf(value interface{}) int {
	var index int = -1
	if list.size == 0 {
		index = -1
	}
	for i, element := range list.elements {
		if reflect.DeepEqual(element, value) {
			index = i
		}
	}
	return index
}

//listIterator()
//实现

//listIterator(int index)
//实现

//remove(int index)
func (list *ArrayList) RemoveWithIndex(index int) error {
	if !list.withinRange(index) {
		return errors.New("out of range")
	} else {
		list.elements = append(list.elements[:index], list.elements[index+1:]...)
		list.size = list.size - 1
		return nil
	}
}

//remove(Object o)
func (list *ArrayList) RemoveWithObeject(object interface{}) bool {
	index := list.IndexOf(object)
	if index == -1 {
		return false
	} else {
		list.elements = append(list.elements[:index], list.elements[index+1:]...)
		list.size = list.size - 1
		return true
	}

}

//removeAll(Collection<?> c)
//实现

//removeIf(Predicate<? super E> filter)
//实现

//removeRange(int fromIndex, int toIndex)
//this function is portected
//func (list *ArrayList) RemoveRange(fromIndex,toIndex int) error{
//	if !list.withinRange(fromIndex)&&!list.withinRange(toIndex)&&fromIndex>=toIndex {
//		return errors.New("out of range")
//	}else {
//		list.elements=append(list.elements[:fromIndex],list.elements[toIndex+1:]...)
//		list.size=list.size-(toIndex-fromIndex)
//		return nil
//	}
//}

//replaceAll(UnaryOperator<E> operator)
//实现

//retainAll(Collection<?> c)
//实现

//set(int index, E element)
func (list *ArrayList) Set(index int, object interface{}) interface{} {
	if list.withinRange(index) {
		list.elements[index] = object
	} else {
		if index == list.size {
			list.Add(object)
		}
	}
	return object
}

//size()
func (list *ArrayList) Size() int {
	return list.size
}

// sort(Comparator<? super E> c)
func (list *ArrayList) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

//spliterator()
//实现

//subList(int fromIndex, int toIndex)
func (list *ArrayList) SubList(fromIndex, toIndex int) ArrayList {
	if !list.withinRange(fromIndex) && !list.withinRange(toIndex) {
		return ArrayList{}
	} else {
		subArray := list.elements[fromIndex:toIndex]
		return ArrayList{
			subArray,
			len(subArray),
			cap(subArray),
		}
	}
}

// toArray()
func (list *ArrayList) ToArray() []interface{} {
	newElements := make([]interface{}, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

//toArray(T[] a)
//实现
//trimToSize()
func (list *ArrayList) TrimToSize() {
	length := len(list.elements)
	newElements := make([]interface{}, length, length)
	copy(newElements, list.elements)
	//for index,value:= range list.elements{
	//	newElements[index]=value
	//}
	list.cap = length
	list.size = length
	list.elements = newElements
}

// Swap swaps the two values at the specified positions.
func (list *ArrayList) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// String returns a string representation of container
func (list *ArrayList) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *ArrayList) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *ArrayList) reset(newLen, newCapacity int) {
	newElements := make([]interface{}, newLen, newCapacity)
	list.cap = newCapacity
	list.size = newLen
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *ArrayList) growBy(length, capacity int) {
	var newCapacity int
	var newLen int
	currentLen := len(list.elements)
	currentCapacity := cap(list.elements)
	if list.cap+capacity >= currentCapacity {
		newLen = currentLen + length
		newCapacity = int(growthFactor * float32(currentCapacity+capacity))
	} else {
		if list.size+length >= currentLen {
			newLen = currentLen + length
			newCapacity = list.cap
		}
	}
	list.reset(newLen, newCapacity)
}
