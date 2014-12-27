package main

import (
	"github.com/dmelani/data_logger/devices"
	"log"
	"fmt"
)

func main() {
	adxl, err := devices.Devices["adxl345"](0x53, 1)
	if err != nil {
		log.Fatal(err)
	}
	itg, err := devices.Devices["itg3200"](0x68, 1)
	if err != nil {
		log.Fatal(err)
	}

	adxl.Init()
	itg.Init()

	for {
		measurement := itg.Read()
		switch measurement := measurement.(type) {
		case *devices.Acceleration:
			derp := measurement.Value()
			values := derp.([3]int32)
			fmt.Println("Acc! ", values[0], values[1], values[2])
		case *devices.MagneticField:
			fmt.Println("Mag field!")
		case *devices.Gyro:
			derp := measurement.Value()
			values := derp.([3]int32)
			fmt.Println("Gyro! ", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		default:
			fmt.Println("Unknown type")
		}
	}
}
