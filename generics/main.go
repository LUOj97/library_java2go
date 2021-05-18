package generics

import (
	"fmt"
	"reflect"
)

//implements generics through a empty interface

func TestInterface(v interface{}) interface{} {
	switch v.(type) {
	case int:
		// DO Something
		return v.(int) + 10
	case float64:
		// DO Something
		return v.(float64) + 22.3
	}
	return v
}

func main() {
	t1 := TestInterface(10)
	t2 := TestInterface(10.0)
	fmt.Println(t1, reflect.TypeOf(t1).String())
	fmt.Println(t2, reflect.TypeOf(t2).String())
}
