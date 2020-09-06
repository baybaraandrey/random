package main

import (
	"fmt"
	"net"
)

func main() {
	ones, bits := 24, 32
	mask := net.CIDRMask(ones, bits)
	ip := net.ParseIP("192.168.0.1")
	network := ip.Mask(mask)

	fmt.Println("IP: ", ip.String())
	fmt.Println("NETWORK: ", network.String())

}

