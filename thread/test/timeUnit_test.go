package test

import (
	"fmt"
	"library_java2go/thread/TimeUnit"
	"sync"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	x := TimeUnit.Convert(TimeUnit.HOURS, 100, TimeUnit.MINUTES)
	fmt.Println(x)
	Y := TimeUnit.Convert(TimeUnit.DAYS, 200, TimeUnit.MINUTES)
	fmt.Println(Y)
	c := TimeUnit.Convert(TimeUnit.SECONDS, 3600, TimeUnit.MINUTES)
	fmt.Println(c)
}

var wg sync.WaitGroup

func TestSleep(t *testing.T) {
	wg.Add(1)
	fmt.Println(time.Now())
	TimeUnit.Sleep(TimeUnit.SECONDS, 10)
	fmt.Println(time.Now())
}
