package main

import (
	"bufio"
	"fmt"
	"library_java2go/main/codec"
	"net"
	"time"
)

//var quitSemaphore chan bool
//func main()  {
//	socket,_:=socket.NewWithHostPort("localhost",9999)
//
//	defer socket.Close() //关闭连接
//	fmt.Println("Connected!")
//	go onMessageRecived(socket) //接收消息
//	go sendMessage(socket) //发送消息
//	<-quitSemaphore
//}
//func sendMessage(conn *socket.Socket)  {
//	//发送消息
//	for{
//		time.Sleep(1 * time.Second)
//		var msg string
//		fmt.Scanln(&msg)
//		if msg == "quit"{
//			quitSemaphore <- true
//			break
//		}
//		//lk
//		b :=[]byte(msg +"\n")
//		//处理加密
//		//b ,_ := codec.Encode(msg+"\n")
//		conn.TcpConn.Write(b)
//	}
//}
//
//func onMessageRecived(conn *socket.Socket) {
//	reader := bufio.NewReader(conn.TcpConn)
//	for {
//		//解密
//		//msg, err := codec.Decode(reader) //
//		msg, err := reader.ReadString('\n')
//		fmt.Println(msg)
//		if err != nil {
//			quitSemaphore <- true
//			break
//		}
//	}
//}

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr) //开启连接
	defer conn.Close()                          //关闭连接
	fmt.Println("Connected!")
	go onMessageRecived(conn) //接收消息
	go sendMessage(conn)      //发送消息
	<-quitSemaphore
}

// 发送消息
func sendMessage(conn *net.TCPConn) {
	//发送消息
	for {
		time.Sleep(1 * time.Second)
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" {
			quitSemaphore <- true
			break
		}
		//lk
		//b :=[]byte(msg +"\n")
		//处理加密
		b, _ := codec.Encode(msg + "\n")
		conn.Write(b)
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		//解密
		msg, err := codec.Decode(reader) //reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
