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
	"fmt"
	"github.com/emirpasic/gods/utils"
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
		elements: make([]interface{}, 10),
		size:     0,
		cap:      10,
	}
	return list
}

// ArrayList(int initialCapacity)
func NewWithInitialCapacity(capacity int) *ArrayList {
	list := &ArrayList{
		elements: make([]interface{}, capacity),
		size:     0,
		cap:      capacity,
	}
	return list
}

// Add()
func (list *ArrayList) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}

}

//add(int index, E element)
func (list *ArrayList) AddWithIndex(index int, values ...interface{}) {
	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

//addAll(int index, Collection<? extends E> c)
func (list *ArrayList) AddAll(e *ArrayList) {
	for i := 0; i < e.size; i++ {
		list.Add(e.elements[i])
	}
}

//clear()
func (list *ArrayList) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

//contains(Object o)
func (list *ArrayList) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

//ensureCapacity(int minCapacity)
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
		if element == value {
			return index
		}
	}
	return -1
}

//isEmpty()
func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil                                    // cleanup reference
	copy(list.elements[index:], list.elements[index+1:list.size]) // shift to the left by one (slow operation, need ways to optimize this)
	list.size--

	list.shrink()
}

// Values returns all elements in the list.
func (list *ArrayList) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Size returns number of elements within the list.
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

// Swap swaps the two values at the specified positions.
func (list *ArrayList) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *ArrayList) Set(index int, value interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
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

func (list *ArrayList) TrimToSize() {

}

// Check that the index is within bounds of the list
func (list *ArrayList) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap)
	list.cap = cap
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *ArrayList) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *ArrayList) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
