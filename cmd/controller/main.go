package main

import (
	"flag"
	"log"
	"net"
	"reflect"
	"time"

	libvirt "github.com/digitalocean/go-libvirt"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/dippynark/kube-qemu/pkg/client"
	factory "github.com/dippynark/kube-qemu/pkg/informers/externalversions"
)

var (
	// apiserverURL is the URL of the API server to connect to
	apiserverURL = flag.String("apiserver", "http://127.0.0.1:8001", "URL used to access the Kubernetes API server")

	// libvirt connection parameters
	network = flag.String("network", "unix", "tcp, unix")
	address = flag.String("address", "/var/run/libvirt/libvirt-sock", "Connection address")

	// queue is a queue of resources to be processed. It performs exponential
	// backoff rate limiting, with a minimum retry period of 5 seconds and a
	// maximum of 1 minute.
	queue = workqueue.NewRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(time.Second*5, time.Minute))

	// stopCh can be used to stop all the informer, as well as control loops within the application.
	stopCh = make(chan struct{})

	// sharedFactory is a shared informer factory that is used a a cache for
	// items in the API server. It saves each informer listing and watching the
	// same resources independently of each other, thus providing more up to
	// date results with less 'effort'
	sharedFactory factory.SharedInformerFactory

	// cl is a Kubernetes API client for our custom resource definition type
	cl client.Interface
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

	log.Printf("Created QEMU client")

	// create an instance of our own API client
	// rest.Config holds the common attributes for Kubernetes client initialisation
	cl, err := client.NewForConfig(&rest.Config{
		Host: *apiserverURL,
	})
	if err != nil {
		log.Fatalf("error creating api client: %s", err.Error())
	}
	log.Printf("Created Kubernetes client")

	// we use a shared informer from the informer factory, to save calls to the
	// API as we grow our application and so state is consistent between our
	// control loops. We set a resync period of 30 seconds, in case any
	// create/replace/update/delete operations are missed when watching
	sharedFactory = factory.NewSharedInformerFactory(cl, time.Second*30)

	informer := sharedFactory.Hypervisor().V1alpha1().VirtualMachines().Informer()
	// we add a new event handler, watching for changes to API resources.
	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: enqueue,
			UpdateFunc: func(old, cur interface{}) {
				if !reflect.DeepEqual(old, cur) {
					enqueue(cur)
				}
			},
			DeleteFunc: enqueue,
		},
	)
}
