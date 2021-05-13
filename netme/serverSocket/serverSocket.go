package serverSocket

import (
	"net"
)

type ServerSocket struct {
	net.TCPListener
}

//accept() 返回值不对
func (s *ServerSocket) Accept() *net.TCPConn {
	server, _ := s.AcceptTCP()
	return server
}

//bind(SocketAddress endpoint)
//bind(SocketAddress endpoint, int backlog)

//void	close()
func (s *ServerSocket) CloseJ() {
	s.Close()
}

//ServerSocketChannel	getChannel()
//InetAddress	getInetAddress()
func (s *ServerSocket) GetInetAddress() net.Addr {
	return s.Addr()
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
