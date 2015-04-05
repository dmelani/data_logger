package main

import (
	"fmt"
	"github.com/dmelani/data_logger/devices"
	"math"
)

type Ahrs struct {
	pitch float64
	yaw float64
	roll float64
	GforceTotal float64
}

var AhrsDebug bool

func (ahrs *Ahrs) AddMeasurement(m devices.Measurement) {
	switch measurement := m.(type) {
	case *devices.Acceleration:
		value := measurement.Value()
		values := value.([3]int32)

		if AhrsDebug {
			fmt.Println("Acc:", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		}

		ahrs.GforceTotal = math.Sqrt(float64(values[0])/1000.0*float64(values[0])/1000.0 + float64(values[1])/1000.0*float64(values[1])/1000.0 + float64(values[2])/1000.0*float64(values[2])/1000.0)

	case *devices.MagneticField:
		value := measurement.Value()
		values := value.([3]int32)

		if AhrsDebug {
			fmt.Println("Mag:", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		}

	case *devices.Gyro:
		value := measurement.Value()
		values := value.([3]int32)

		if AhrsDebug {
			fmt.Println("Gyro:", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		}

	default:
		fmt.Println("Unknown type")
	}
}
