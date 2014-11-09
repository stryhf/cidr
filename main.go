//
// cidr: display IP range of specified CIDR block
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
		b[i] = n.IP[i] + (0xff ^ n.Mask[i])
	}
	return b
}

func main() {
	p := fmt.Println
	
	if len(os.Args) < 2 {
		p("cidr: missing network")
		return
	}

	_, net, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		p(err)
		return
	}
	fmt.Println(" network:", net.IP, netCast(net))
}
