package main

import (
	"fmt"
	"net"
)

func main() {
	ones, bits := 24, 32
	mask := net.CIDRMask(ones, bits)
	ip, _ := net.ParseIP("192.168.0.1")
	network := ip.Mask(mask)

	fmt.Println(ip.String())
	fmt.Println(network.String())

}

