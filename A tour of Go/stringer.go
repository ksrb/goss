//http://tour.golang.org/methods/7
package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (addr IPAddr) String2() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

// Initial attempt
func (addr IPAddr) String() string {
	str := ""
	for index, value := range addr {
		str += fmt.Sprintf("%d", value)
		if index != len(addr)-1 {
			str += "."
		}
	}
	return str
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
