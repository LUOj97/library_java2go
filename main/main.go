package main

import (
	"bufio"
	"fmt"
	"library_java2go/main/codec"
	"net"
)

//func main() {
//	name, err := os.Hostname()
//	if err != nil {
//		fmt.Printf("Oops: %v\n", err)
//		return
//	}
//
//	addrs, err := net.LookupHost(name)
//	if err != nil {
//		fmt.Printf("Oops: %v\n", err)
//		return
//	}
//
//	for _, a := range addrs {
//		fmt.Println(a)
//	}
//}
//var q SynchronousQueue.SynchronousQueue
//var wg 	sync.WaitGroup
//func main()  {
// 	wg.Add(3)
//	fmt.Println("test")
//	go one()
//	go two()
//	go three()
//	defer wg.Wait()
//
//}
//func one() {
//	defer wg.Done()
//	fmt.Println("this is therea 1")
//	q.Push(1)
//	fmt.Println("1 push end ")
//	time.Sleep(10)
//}
//
//func two() {
//	defer wg.Done()
//	fmt.Println("this is thread 2")
//	q.Push(2)
//	q.Push(3)
//	q.Push(4)
//	fmt.Println("2 push end ")
//}
//func three() {
//	defer wg.Done()
//	fmt.Println("this is test 3")
//	x:=q.Pop()
//	fmt.Println("TAKE ",x)
//}

//func main()  {
//	a:=time.Duration(1000000000)
//	fmt.Println(a.Seconds())
//
//}

var ConnMap map[string]*net.TCPConn

func main() {
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPConn) //初始化
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:9999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr) //开启tcp 服务
	//退出时关闭
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		// 新连接加入 map
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn

		go tcpPipe(tcpConn)
	}
}

//处理发送过来的消息
func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected : " + ipStr)
		conn.Close()
	}()
	//读取数据
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(string(message))
		//这里返回消息改为广播
		boradcastMessage(conn.RemoteAddr().String() + ":" + string(message))
	}
}

//广播给其它
func boradcastMessage(message string) {
	//遍历所有客户端并发消息
	for _, conn := range ConnMap {
		b, err := codec.Encode(message)
		if err != nil {
			continue
		}
		conn.Write(b)
	}
}
