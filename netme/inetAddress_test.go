package netme

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {
	addr, _ := GetAddress("www.baidu.com")
	fmt.Println(addr.String())
}
