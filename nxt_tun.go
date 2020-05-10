package main

import (
	“log”
	“github.com/songgao/packets/ethernet”
	“github.com/songgao/water”
)

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}

	config.Name = “sample_dev”

	ifce, err := water.New(config) //This is where the tap device is created
	if err != nil {
		log.Fatal(err)
	}

	var frame ethernet.Frame
	for {
		frame.Resize(1500)

		n, err := ifce.Read([]byte(frame)) //the userspace program now reads

		//on the file descriptor
		if err != nil {
			log.Fatal(err)
		}

		//all the details of the ethernet frame are printed

		frame = frame[:n]

		log.Printf(“Dst: %s\n”, frame.Destination())
		log.Printf(“Src: %s\n”, frame.Source())
		log.Printf(“Ethertype: % x\n”, frame.Ethertype())
		log.Printf(“Payload: % x\n”, frame.Payload())
	}
}