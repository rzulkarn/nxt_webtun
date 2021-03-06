package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/songgao/water"
	"golang.org/x/net/ipv4"
)

var (
	snaplen int32 = 65535
	MTU     int64 = 1500
)

func ifaceSetup(localCIDR string) *water.Interface {
	iface, err := water.New(water.Config{DeviceType: water.TUN})
	if nil != err {
		log.Println("Unable to allocate TUN interface:", err)
		panic(err)
	}

	log.Println("Interface name allocated:", iface.Name())

	if err := exec.Command("ifconfig", iface.Name(), "inet", "10.1.0.1", localCIDR, "mtu", strconv.FormatInt(MTU, 10), "up").Run(); err != nil {
		log.Fatalln("Unable to setup interface:", err)
	}
	return iface
}

func main() {
	ifce := ifaceSetup("10.1.0.2")

	packet := make([]byte, snaplen)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Packet Received %d: % x\n", n, packet[:n])

		// Parse the packet header -
		header, _ := ipv4.ParseHeader(packet[:n])
		fmt.Printf("Header struct: %+v (%+v)\n", header, err)

		newPacket := gopacket.NewPacket(packet, layers.LayerTypeIPv4, gopacket.Default)

		// Check if packet is TCP ....
		if tcpLayer := newPacket.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			fmt.Println("This is TCP packet!")
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Printf("From src port %d to dst port %d\n", tcp.SrcPort, tcp.DstPort)
		}

		// Check if packet is IPv4 ....
		if ipLayer := newPacket.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			//fmt.Println("This is IP packet")
			ip, _ := ipLayer.(*layers.IPv4)
			//fmt.Printf("IP Dst: %s, Src: %s, Payload: % x\n", ip.DstIP.String(), ip.SrcIP.String(), ip.Payload)

			ip.DstIP = net.ParseIP("127.0.0.1")

			options := gopacket.SerializeOptions{
				ComputeChecksums: true,
				FixLengths:       true,
			}

			newBuffer := gopacket.NewSerializeBuffer()
			err := gopacket.SerializePacket(newBuffer, options, newPacket)
			if err != nil {
				log.Printf("[-] Serialize error: %s\n", err.Error())
				return
			}

			outgoingPacket := newBuffer.Bytes()
			fmt.Printf("Outgoing Packet % x\n", outgoingPacket)
			fmt.Println("====")

			// write it back into utun interface
			ifce.Write(outgoingPacket)
		} else {
			log.Printf("[-] Not an IPv4 packet\n")
		}
	}
}

//  if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
// 	tcp := tcpLayer.(*layers.TCP)
// 	dst = fmt.Sprintf("%s:%d", dst, tcp.DstPort)
// 	src = fmt.Sprintf("%s:%d", src, tcp.SrcPort)
// 	fmt.Printf("From %s to %s\n\n", src, dst)
// 	ip.DstIP = net.ParseIP("8.8.8.8")
// 	newBuffer := gopacket.NewSerializeBuffer()
// 	gopacket.SerializeLayers(newBuffer, options,
// 		ip,
// 		tcp,
// 	)
// var frame ethernet.Frame
// for {
// 	frame.Resize(1500)

// nc
