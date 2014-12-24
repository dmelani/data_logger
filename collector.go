package main

import (
	"github.com/dmelani/telemetry_collector/devices"
	"log"
	"fmt"
)

func main() {
	adxl, err := devices.Devices["adxl345"](0x53, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(adxl)
	adxl.Init()
	for {
		fmt.Println(adxl.Read())
	}
}
