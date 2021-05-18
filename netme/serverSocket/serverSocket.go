package serverSocket

import (
	"errors"
	"fmt"
	"library_java2go/netme/socket"
	"net"
	"time"
)

type ServerSocket struct {
	TcpListen *net.TCPListener
}

//ServerSocket()
func New() *ServerSocket {
	serverSocket := new(ServerSocket)
	return serverSocket
}

//ServerSocket(int port)
func NewWithPort(port int) (*ServerSocket, error) {
	server := fmt.Sprintf(":%d", port)
	addr, _ := net.ResolveTCPAddr("tcp", server)
	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		e := errors.New("connect error")
		return nil, e
	}
	serverSocket := ServerSocket{tcp}
	return &serverSocket, nil
}

//ServerSocket(int port, int backlog) backlog set before
//func NewWithPortBacklog(port,backlog int)  (*ServerSocket,error){
//	server := fmt.Sprintf(":%d",port)
//	addr, _ := net.ResolveTCPAddr("tcp", server)
//	tcp, err := net.ListenTCP("tcp", addr)
//	if err != nil {
//		e := errors.New("connect error")
//		return nil, e
//	}
//	serverSocket := ServerSocket{tcp}
//	return &serverSocket, nil
//}
//ServerSocket(int port, int backlog, InetAddress bindAddr)

//accept()
func (s *ServerSocket) Accept() *socket.Socket {
	server, _ := s.TcpListen.AcceptTCP()
	return &socket.Socket{TcpConn: server}
}

//bind(SocketAddress endpoint)
func (s *ServerSocket) Bind(addr net.TCPAddr) error {
	tcp, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		e := errors.New("connect error")
		return e
	}
	s.TcpListen = tcp
	return nil
}

//bind(SocketAddress endpoint, int backlog)

//void	close()
func (s *ServerSocket) Close() {
	s.TcpListen.Close()
}

//ServerSocketChannel	getChannel()
//InetAddress	getInetAddress()
func (s *ServerSocket) GetInetAddress() net.Addr {
	return s.TcpListen.Addr()
}

//int	getLocalPort()
//SocketAddress	getLocalSocketAddress()

//int	getReceiveBufferSize()
//boolean	getReuseAddress()
//int	getSoTimeout()
//protected void	implAccept(Socket s)

//boolean	isBound()
//boolean	isClosed()
//setPerformancePreferences(int connectionTime, int latency, int bandwidth)
//setReceiveBufferSize(int size)
//setReuseAddress(boolean on)
//setSocketFactory(SocketImplFactory fac)
//setSoTimeout(int timeout)
func (s *ServerSocket) SetSoTimeout(timeout int) {
	duration := time.Duration(timeout)
	s.TcpListen.SetDeadline(time.Now().Add(duration * time.Millisecond))
}
