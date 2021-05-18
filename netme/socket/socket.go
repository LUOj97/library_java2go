package socket

import (
	"context"
	"errors"
	"fmt"
	"library_java2go/netme/inetAddress"
	"net"
	"time"
)

type Socket struct {
	TcpConn *net.TCPConn
}
type Person struct {
	name string
}

//Socket()
func New() *Socket {
	socket := new(Socket)
	return socket
}

//Socket(InetAddress address, int port)
func NewWithInetAddressHost(inetaddre inetAddress.InetAddress, port int) (*Socket, error) {
	tcpAddress := net.TCPAddr{
		inetaddre.Ip,
		port,
		"",
	}
	tcp, err := net.DialTCP("tcp", nil, &tcpAddress)
	if err != nil {
		e := errors.New("connect error")
		return nil, e
	}
	socekt := Socket{tcp}
	return &socekt, nil
}

//Socket(InetAddress address, int port, InetAddress localAddr, int localPort)
func NewWithLocaladdr(inetaddre, localaddr inetAddress.InetAddress, port, localPort int) (*Socket, error) {
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
	tcp, err := net.DialTCP("tcp", &localAddr, &addre)
	if err != nil {
		e := errors.New("connect error")
		return nil, e
	}
	socekt := Socket{tcp}
	return &socekt, nil
}

//Socket(Proxy proxy)
//Socket(SocketImpl impl) protected

//Socket(String host, int port)
func NewWithHostPort(host string, port int) (*Socket, error) {
	server := fmt.Sprintf("%s:%d", host, port)
	addr, _ := net.ResolveTCPAddr("tcp", server)
	tcp, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		e := errors.New("connect error")
		return nil, e
	}
	socekt := Socket{tcp}
	return &socekt, nil
}

//bind(SocketAddress bindpoint)
func (s *Socket) Bind(addr net.TCPAddr) {
	remoteAddr := s.GetRemoteSocketAddress()
	rAddr := remoteAddr.(*net.TCPAddr)
	x, _ := net.DialTCP("tcp", &addr, rAddr)
	s.TcpConn = x

}

//close()
func (s *Socket) Close() {
	s.TcpConn.Close()
}

//connect(SocketAddress endpoint)
func (s *Socket) Connect(addr net.TCPAddr, timeout int64) {
	x, _ := net.DialTCP("tcp", nil, &addr)
	s.TcpConn = x
}

//connect(SocketAddress endpoint, int timeout)
func (s *Socket) ConnectWithTimeout(addr net.TCPAddr, timeout int64) {
	var d net.Dialer
	duration := time.Duration(timeout)
	ctx, _ := context.WithTimeout(context.Background(), duration)
	address := addr.IP.String()
	port := addr.Port
	address = fmt.Sprintf("%s:%d", address, port)
	tcpconn, _ := d.DialContext(ctx, "tcp", address)
	TcpConn := tcpconn.(*net.TCPConn)
	s.TcpConn = TcpConn
}

//getChannel()
func (s *Socket) GetChannel() *SocketChannel {
	return nil
}

type SocketChannel struct {
}

//getInetAddress()
func (s *Socket) GetInetAddress() net.Addr {
	return s.TcpConn.RemoteAddr()
}

//getInputStream()  write
func (s *Socket) Write(b []byte) (int, error) {
	return s.TcpConn.Write(b)
}

//getOutputStream()  read
func (s *Socket) Read(b []byte) (int, error) {
	return s.TcpConn.Read(b)
}

//getKeepAlive()

//InetAddress	getLocalAddress() 需要修改
func (s *Socket) GetLocalAddress() net.Addr {
	return s.TcpConn.LocalAddr()
}

//getLocalPort()
func (s *Socket) GetLocalPort() int {
	adre := s.TcpConn.LocalAddr()
	port := adre.(*net.TCPAddr).Port
	return port
}

//getLocalSocketAddress()

//getOOBInline()

//getPort()
func (s *Socket) GetPort() int {
	adre := s.GetRemoteSocketAddress()
	port := adre.(*net.TCPAddr).Port
	return port
}

//getReceiveBufferSize()

//getRemoteSocketAddress()
func (s *Socket) GetRemoteSocketAddress() net.Addr {
	addr := s.TcpConn.RemoteAddr()
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
	s.TcpConn.SetKeepAlive(on)
}

//setOOBInline(boolean on)

//setPerformancePreferences(int connectionTime, int latency, int bandwidth)

//setReceiveBufferSize(int size)
func (s *Socket) SetReceiveBufferSize(size int) {
	s.TcpConn.SetReadBuffer(size)
}

//setReuseAddress(boolean on)

//setSendBufferSize(int size)
func (s *Socket) SetSendBufferSize(size int) {
	s.TcpConn.SetWriteBuffer(size)
}

//setSocketImplFactory(SocketImplFactory fac)

//setSoLinger(boolean on, int linger)
func (s *Socket) SetSoLinger(on bool, sec int) {
	s.TcpConn.SetLinger(sec)
}

//setSoTimeout(int timeout)

//setTcpNoDelay(boolean on)
func (s *Socket) SetTcpNoDelay(on bool) {
	s.TcpConn.SetNoDelay(on)
}

//void	setTrafficClass(int tc)
//void	shutdownInput()
func (s *Socket) ShutdownInput() {
	s.TcpConn.CloseWrite()
}

//void	shutdownOutput()
func (s *Socket) ShutdownOutput() {
	s.TcpConn.CloseRead()
}

//string	toString()
