package netme

import "net"

type InetAddress struct {
	net.IP
}

//
func (ip InetAddress) IsMulticastAddress() bool {
	return ip.IsMulticast()
}
