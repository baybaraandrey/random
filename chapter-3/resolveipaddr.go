package main

import (
	"fmt"
	"os"
	"net"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: resolve <hostname>\n")
		os.Exit(1)
	}
	
	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot resolve hostname %s\n", os.Args[1])
		os.Exit(1)
	}
	
	fmt.Fprintf(os.Stdout, "Resolved network is %q\n", addr.Network())
	fmt.Fprintf(os.Stdout, "Resolved address is %q\n", addr.String())
	os.Exit(0)
}


