package main

import (
	"flag"
	"log"
	"net"
	"time"

	libvirt "github.com/digitalocean/go-libvirt"
)

var (
	// libvirt connection parameters
	network = flag.String("network", "unix", "tcp, unix")
	address = flag.String("address", "/var/run/libvirt/libvirt-sock", "Connection address")

	// stopCh can be used to stop all the informer, as well as control loops within the application.
	stopCh = make(chan struct{})
)

func main() {
	flag.Parse()

	// connect to address on named network
	c, err := net.DialTimeout(*network, *address, 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	// configure libvirt RPC connection
	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

}
