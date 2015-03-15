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
		Bus int
		Port int
	}
}

func setupDevices(config Config) (r []devices.Devices) {
	for _, e := range config.Sensors {
		log.Println(e)
		dev, err := devices.Devices[e.Type](e.Port, e.Bus)
		if err != nil {
			log.Fatal(err)
		}
		append(r, dev)
	}
}

func initDevices(devices []devices.Devices) {
	for _, e := range devices {
		e.Init()
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

	devices := setupDevices(cfg)
	initDevices(devices)

	for {
		measurement := itg.Read()
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
