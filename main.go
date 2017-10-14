package main

import (
	"flag"
	"log"
	"net"
	"time"

	"k8s.io/client-go/rest"
	libvirt "github.com/digitalocean/go-libvirt"

	"github.com/dippynark/kube-qemu/pkg/apis/hypervisor/v1alpha1"
	"github.com/dippynark/kube-qemu/pkg/client"
	factory "github.com/dippynark/kube-qemu/pkg/informers/externalversions"
)

var (
	// apiserverURL is the URL of the API server to connect to
	apiserverURL = flag.String("apiserver", "http://127.0.0.1:8001", "URL used to access the Kubernetes API server")

	// libvirt connection parameters
	network = flag.String("network", "unix", "tcp, unix")
	address = flag.String("address", "/var/run/libvirt/libvirt-sock", "Connection address")

	// stopCh can be used to stop all the informer, as well as control loops within the application.
	stopCh = make(chan struct{})
)

func main() {
	flag.Parse()

	// connect to address on named network
	conn, err := net.DialTimeout(*network, *address, 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	// configure libvirt RPC connection
	l := libvirt.New(conn)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// create an instance of our own API client
	// rest.Config holds the common attributes for Kubernetes client initialisation
	cl, err := client.NewForConfig(&rest.Config{
		Host: *apiserverURL
	})
	if err != nil {
		log.Fatalf("error creating api client: %s", err.Error())
	}
	log.Printf("Created Kubernetes client")
}
