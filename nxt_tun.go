package main

import (
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	"log"
)

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}

	config.Name = "sample_dev"

	ifce, err := water.New(config) // New a tap device
	if err != nil {
		log.Fatal(err)
	}

	var frame ethernet.Frame
	for {
		frame.Resize(1500)

		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}

		frame = frame[:n]

		log.Printf("Dst: %s\n", frame.Destination())
	}
}
