package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// devices, _ := pcap.FindAllDevs()

	// for _, device := range devices {
	// 	fmt.Println(device.Name)
	// 	fmt.Println(device.Description)
	// 	fmt.Println(device.Flags)
	// 	fmt.Println(device.Addresses)
	// 	fmt.Println()
	// }
	log.Println("start")
	defer log.Println("end")

	// timeout := time.Duration(30) * time.Second

	// handle, err := pcap.OpenLive("en0", int32(1024), true, timeout)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer handle.Close()

	// packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// for packet := range packetSource.Packets() {
	// 	fmt.Println(packet)
	// }

	if handle, err := pcap.OpenOffline("pcapfiles/pingdup.pcap"); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			// fmt.Println(packet.NetworkLayer().NetworkFlow().Dst())
			if packet == nil {
				fmt.Println(packet)
			}
		}
	}
}
