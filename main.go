package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/vishvananda/netlink"
)

// Load the compiled object file
func loadEBPF() (*ebpf.Collection, error) {
	spec, err := ebpf.LoadCollectionSpec("drop_packets.o")
	if err != nil {
		return nil, err
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

func main() {
	// Parse cli argument
	ifaceName := flag.String("iface", "", "Network interface to attach.")
	flag.Parse()

	if *ifaceName == "" {
		log.Fatal("Usage: go run main.go -iface <interface-name>")
	}

	// Get the interface index
	iface, err := netlink.LinkByName(*ifaceName)
	if err != nil {
		log.Fatalf("Failed to get interface %s: %v", *ifaceName, err)
	}

	// Load
	coll, err := loadEBPF()
	if err != nil {
		log.Fatalf("Failed to load Dropify: %v", err)
	
	defer coll.Close()

	prog := coll.Programs["drop_packets"]
	if prog == nil {
		log.Fatalf("Failed to find.")
	}

	// Attach
	tcHook, err := link.AttachTCX(link.TCOptions{
		Program:   prog,
		Interface: iface.Attrs().Index,
		Attach:    link.AttachTCXIngress,
	})
	if err != nil {
		log.Fatalf("Failed to attach: %v", err)
	}
	defer tcHook.Close()

	fmt.Printf("Dropify attached to %s, dropping all packets.\n", *ifaceName)

	// Wait for a termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Detaching Dropify and exiting.")
}
