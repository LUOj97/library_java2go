package socket

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestNew(t *testing.T) {
	x := New()
	fmt.Println(x)
}
func TestNewWithHostPort(t *testing.T) {
	//c,err:=NewWithHostPort("localhost",8888)
	//if err != nil {
	//	log.Println("dial error:", err)
	//	return
	//}
	//defer c.Close()
	//log.Println("dial ok")
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
}
func TestName(t *testing.T) {

}
