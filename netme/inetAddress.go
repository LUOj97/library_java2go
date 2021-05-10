package netme

import "net"

type InetAddress []byte

//getAddress()
//func (inetAddr *InetAddress) GetAddress() net.IP {
//	//return inetAddr.ip
//
//}

//static
//getAllByName(String host) , static InetAddress[]
func GetAllByName(host string) ([]InetAddress, error) {
	iprecords, _ := net.LookupIP(host)
	ips := make([]InetAddress, len(iprecords))
	for i, ip := range iprecords {
		ipN := ip
		ips[i] = InetAddress(ipN)
	}
	return ips, nil
}

//static
//getByAddress(byte[] addr) ,static InetAddress

//static
//getByAddress(String host, byte[] addr),static InetAddress

//static
//getByName(String host) static InetAddress
