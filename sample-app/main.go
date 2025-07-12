package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	ans := ""
	for i := 0; i < len(ip); i++ {
		if i > 0 {
			ans += "."
		}
		ans += fmt.Sprintf("%d", ip[i])
	}
	return ans
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
