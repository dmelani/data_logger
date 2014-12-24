package main

import (
	"github.com/dmelani/telemetry_collector/devices"
	"log"
)

func main() {
	adxl, err := devices.NewAdxl345(0x53, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(adxl)
	adxl.Init()
}

