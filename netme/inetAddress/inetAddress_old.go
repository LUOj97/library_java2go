package inetAddress

import (
	"context"
	"net"
	"os"
)

type InetAddress struct {
	ip   net.IP
	host string
}

////static
//getAllByName(String host) , static InetAddress[]
func GetAllByName(host string) ([]InetAddress, error) {
	var DefaultResolver = &net.Resolver{}
	addrs, err := DefaultResolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return nil, err
	}
	ips := make([]InetAddress, len(addrs))
	for i, ia := range addrs {
		inetaddr := InetAddress{
			ia.IP,
			host,
		}
		ips[i] = inetaddr
	}
	return ips, nil
}

//static
//getByAddress(byte[] addr) ,static InetAddress

//static
//getByAddress(String host, byte[] addr),static InetAddress

//static
//getByName(String host) static InetAddress
func GetByName(host string) (InetAddress, error) {
	x, _ := GetAllByName(host)
	return x[0], nil
}

//getCanonicalHostName()

//getHostAddress()
func (ip *InetAddress) GetHostAddress() string {
	return ip.ip.String()
}

//getHostName()
func (ip *InetAddress) GetHostName() string {
	return ip.host
}

//static
//getLocalHost()
func GetLocalHost() (*InetAddress, error) {
	name, err1 := os.Hostname()
	addrs, err2 := net.LookupHost(name)
	if err1 != nil && err2 != nil {
		return nil, nil
	}
	localIP := net.ParseIP(addrs[0])
	return &InetAddress{localIP, name}, nil
}

//static
//getLoopbackAddress()
func GetLoopbackAddress() (*InetAddress, error) {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip.IsLoopback() {
			return &InetAddress{ip, ""}, nil
		}

	}

	return nil, nil
}

//hashCode()

//isAnyLocalAddress()
func (ip InetAddress) IsAnyLocalAddress() bool {
	return false
}

//isLinkLocalAddress()
func (ip InetAddress) IsLinkLocalAddress() bool {
	//return ip.ip.IsLinkLocalMulticast() || ip.ip.IsLinkLocalUnicast()
	return false
}

//isLoopbackAddress()
func (ip InetAddress) IsLoopbackAddress() bool {
	//return ip.ip.IsLoopback()
	return false
}

//isMCGlobal()
func (ip InetAddress) IsMCGlobal() bool {
	return false
}

//isMCLinkLocal()
func (ip InetAddress) IsMCLinkLocal() bool {
	return false
}

//isMCNodeLocal()
func (ip InetAddress) IsMCNodeLocal() bool {
	return false
}

//isMCOrgLocal()
func (ip InetAddress) IsMCOrgLocal() bool {
	return false
}

//isMCSiteLocal()
func (ip InetAddress) IsMCSiteLocal() bool {
	return false
}

//isMulticastAddress()
func (ip InetAddress) IsMulticastAddres() bool {
	//return ip.ip.IsMulticast()
	return false
}

//isReachable(int timeout)

//isReachable(NetworkInterface netif, int ttl, int timeout)

//isSiteLocalAddress()

//toString()
func (ip InetAddress) ToString() {
	ip.ip.String()
}
