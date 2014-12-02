//
// cidr: display network, netmask and broadcast of specified CIDR block
//

package main

import (
	"fmt"
	"net"
	"os"
)

func netCast(n *net.IPNet) net.IP {
	b := make(net.IP, net.IPv4len)

	for i := 0; i < 4; i++ {
		b[i] = n.IP[i] + ^n.Mask[i]
	}
	return b
}

func main() {
	p := fmt.Println

	if len(os.Args) < 2 {
		p("Usage: cidr <NETWORK>")
		return
	}

	_, net, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		p(err)
		return
	}
	fmt.Printf(" network: %-15s / %s\n          %s\n", net.IP, net.Mask, netCast(net))
}
