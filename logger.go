package main

import (
	"github.com/dmelani/data_logger/devices"
	"log"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
)

type Config struct {
	Sensors []struct {
		Name string
		Type string
	}
}

func main() {
	configFile := os.Args[1]
	cfg := Config{}

	configData, err := ioutil.ReadFile(configFile);
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	adxl, err := devices.Devices["adxl345"](0x53, 1)
	if err != nil {
		log.Fatal(err)
	}
	itg, err := devices.Devices["itg3200"](0x68, 1)
	if err != nil {
		log.Fatal(err)
	}
	hmc, err := devices.Devices["hmc5883l"](0x1E, 1)
	if err != nil {
		log.Fatal(err)
	}

	adxl.Init()
	itg.Init()
	hmc.Init()

	for {
		measurement := hmc.Read()
		switch measurement := measurement.(type) {
		case *devices.Acceleration:
			derp := measurement.Value()
			values := derp.([3]int32)
			fmt.Println("Acc:", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		case *devices.MagneticField:
			derp := measurement.Value()
			values := derp.([3]int32)
			fmt.Println("Mag:", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		case *devices.Gyro:
			derp := measurement.Value()
			values := derp.([3]int32)
			fmt.Println("Gyro:", float32(values[0])/1000.0, float32(values[1])/1000.0, float32(values[2])/1000.0)
		default:
			fmt.Println("Unknown type")
		}
	}
}
