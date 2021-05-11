package inetAddress

import (
	"fmt"
	"testing"
)

func TestGetAllByName(t *testing.T) {
	//x, _ := GetAllByName("landv.cn")
	//fmt.Println(x)
	//for _, i := range x {
	//	fmt.Println(i.ip)
	//}
	//
	//y,_:=GetByName("www.baidu.com") fmt.Println(y)
	y, _ := GetLoopbackAddress()
	fmt.Println(y)
}
