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
		measurement := adxl.Read()
		switch measurement := measurement.(type) {
		case *devices.Acceleration:
			derp := measurement.Value()
			values := derp.([3]int32)
			fmt.Println("Acc! ", values[0], values[1], values[2])
		case *devices.MagneticField:
			fmt.Println("Mag field!")
		case *devices.Gyro:
			fmt.Println("Gyro!")
		default:
			fmt.Println("Unknown type")
		}
	}
}
