package inetAddress

import (
	"fmt"
	"net"
	"testing"
)

func TestGetAllByName(t *testing.T) {
	x, _ := GetAllByName("landv.cn")
	fmt.Println(x)
	for _, i := range x {
		fmt.Println(i.Ip)
	}

}

func TestGetByName(t *testing.T) {
	y, _ := GetByName("www.baidu.com")
	fmt.Println(y)
}

func TestGetByAddress(t *testing.T) {
	addr := net.ParseIP("104.193.88.123")
	y, _ := GetByAddress(addr)
	fmt.Println(y)
}

func TestGetByAddressWithHost(t *testing.T) {
	addr := net.ParseIP("104.193.88.123")
	host := "baidu.com"
	y, _ := GetByAddressWithHost(host, addr)
	fmt.Println(y)
}

func TestGetLocalHost(t *testing.T) {
	x, _ := GetLocalHost()
	fmt.Println(x)
}

func TestGetLoopbackAddress(t *testing.T) {
	x, _ := GetLoopbackAddress()
	fmt.Println(x)
}
