package netme

import "net"

type InetAddress struct {
	net.IP
}

//getAddress()
func (inetAddr InetAddress) GetAddress() InetAddress {
	return inetAddr
}

//getAllByName(String host)
func GetAllByName(addr string) (*net.IPAddr, error) {
	return net.ResolveIPAddr("ip", addr)
}
