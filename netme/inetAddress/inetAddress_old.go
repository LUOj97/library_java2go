package inetAddress

import (
	"context"
	"net"
)

type InetAddress struct {
	ip   net.IP
	host string
}

//
//getAddress()
//func (inetAddr *InetAddress) GetAddress() net.IP {
//	//return inetAddr.ip
//
//}

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
	//	iprecords, ok:=net.LookupIP(host)
	//	//fmt.Println(iprecords)
	//	if ok==nil {
	//		ips := make([]InetAddress, len(iprecords))
	//		for i,ip:=range iprecords{
	//				inetaddr := InetAddress{
	//					net.IP(ip),
	//					host,
	//				}
	//				//fmt.Println(ip)
	//				//fmt.Println()
	//			fmt.Println(reflect.TypeOf(inetaddr.ip))
	//			fmt.Println(reflect.TypeOf(ip))
	//			fmt.Println(inetaddr)
	//			fmt.Println(inetaddr.ip)
	//				//fmt.Println(net.IP(ip))
	//				ips[i] = inetaddr
	//
	//		}
	//		return ips,nil
	//}
	//return nil,nil
}

//static
//getByAddress(byte[] addr) ,static InetAddress

//static
//getByAddress(String host, byte[] addr),static InetAddress

//static
//getByName(String host) static InetAddress

//getCanonicalHostName()
//getHostAddress()
//getHostName()

//static
//getLocalHost()
//static
//getLoopbackAddress()

//hashCode()
//isAnyLocalAddress()
//isLinkLocalAddress()
func (ip InetAddress) IsLinkLocalAddress() bool {
	return ip.ip.IsLinkLocalMulticast() || ip.ip.IsLinkLocalUnicast()
}

//isLoopbackAddress()
func (ip InetAddress) IsLoopbackAddress() bool {
	return ip.ip.IsLoopback()
}

//isMCGlobal()
//isMCLinkLocal()
//isMCNodeLocal()
//isMCOrgLocal()
//isMCSiteLocal()
//isMulticastAddress()
func (ip InetAddress) IsMulticastAddres() bool {
	return ip.ip.IsMulticast()

}

//isReachable(int timeout)
//isReachable(NetworkInterface netif, int ttl, int timeout)
//isSiteLocalAddress()
//toString()
func (ip InetAddress) ToString() {
	ip.ip.String()
}
