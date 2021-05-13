package socket

import (
	"fmt"
	"library_java2go/netme/inetAddress"
	"net"
)

type Socket struct {
	socket net.TCPConn
}
type Person struct {
	name string
}

//Socket()

//Socket(InetAddress address, int port)
func NewWithInetAddressHost(inetaddre inetAddress.InetAddress, port int) (*net.TCPConn, error) {
	tcpAddress := net.TCPAddr{
		inetaddre.Ip,
		port,
		"",
	}
	return net.DialTCP("tcp", nil, &tcpAddress)
}

//Socket(InetAddress address, int port, InetAddress localAddr, int localPort)
func NewWithLocaladdr(inetaddre, localaddr inetAddress.InetAddress, port, localPort int) (*net.TCPConn, error) {
	addre := net.TCPAddr{
		inetaddre.Ip,
		port,
		"",
	}
	localAddr := net.TCPAddr{
		localaddr.Ip,
		localPort,
		"",
	}
	return net.DialTCP("tcp", &localAddr, &addre)
}

//Socket(Proxy proxy)
//Socket(SocketImpl impl) protected

//Socket(String host, int port)
func NewWithHostPort(host string, port int) (*net.TCPConn, error) {
	server := fmt.Sprintf("%s:%d", host, port)
	addr, _ := net.ResolveTCPAddr("tcp", server)
	return net.DialTCP("tcp", nil, addr)
}

//bind(SocketAddress bindpoint)
//close()
func (s *Socket) Close() {
	s.socket.Close()
}

//connect(SocketAddress endpoint)
//connect(SocketAddress endpoint, int timeout)
//getChannel()
//getInetAddress()
//getInputStream()
//getKeepAlive()

//InetAddress	getLocalAddress() 需要修改
func (s *Socket) GetLocalAddress() {
	addr := s.socket.LocalAddr()
	fmt.Println(addr)
}

//getLocalPort()
//getLocalSocketAddress()
//getOOBInline()
//getOutputStream()
//getPort()
//getReceiveBufferSize()

//getRemoteSocketAddress()
func (s *Socket) GetRemoteSocketAddress() net.Addr {
	addr := s.socket.RemoteAddr()
	return addr
}

//getReuseAddress()
//getSendBufferSize()
//getSoLinger()
//getSoTimeout()
//getTcpNoDelay()
//getTrafficClass()
//isBound()
//isClosed()
//isConnected()
//isInputShutdown()
//isOutputShutdown()
//sendUrgentData(int data)
//setKeepAlive(boolean on)
func (s *Socket) SetKeepAlive(on bool) {
	s.socket.SetKeepAlive(on)
}

//setOOBInline(boolean on)

//setPerformancePreferences(int connectionTime, int latency, int bandwidth)

//setReceiveBufferSize(int size)
func (s *Socket) SetReceiveBufferSize(size int) {
	s.socket.SetReadBuffer(size)
}

//setReuseAddress(boolean on)

//setSendBufferSize(int size)
func (s *Socket) SetSendBufferSize(size int) {
	s.socket.SetWriteBuffer(size)
}

//setSocketImplFactory(SocketImplFactory fac)

//setSoLinger(boolean on, int linger)
func (s *Socket) SetSoLinger(on bool, sec int) {
	s.socket.SetLinger(sec)
}

//setSoTimeout(int timeout)
//setTcpNoDelay(boolean on)
func (s *Socket) SetTcpNoDelay(on bool) {
	s.socket.SetNoDelay(on)
}

//void	setTrafficClass(int tc)
//void	shutdownInput()
//void	shutdownOutput()
//string	toString()
