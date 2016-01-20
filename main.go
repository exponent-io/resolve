package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "invalid arguments %v\nusage: resolve <hostname>\n", args)
		os.Exit(1)
	}
	ips, err := net.LookupHost(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
	if len(ips) == 0 {
		fmt.Fprintf(os.Stderr, "host %q not found", args[0])
		os.Exit(1)
	}
	fmt.Printf("%v", ips[0])
}
