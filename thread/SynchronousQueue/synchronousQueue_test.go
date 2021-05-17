package SynchronousQueue

import (
	"fmt"
	"testing"
	"time"
)

var q SynchronousQueue

func TestSynchronousQueue_Push(t *testing.T) {

	go func() {
		fmt.Println("this is therea 1")
		q.Push(1)
		fmt.Println("1 push end ")
		time.Sleep(10)
	}()

	go func() {
		fmt.Println("this is thread 2")
		q.Push(2)
		q.Push(3)
		q.Push(4)
		fmt.Println("2 push end ")
	}()

	go func() {
		fmt.Println("this is test 3")
		x := q.Pop()
		fmt.Println("TAKE ", x)
	}()

}
