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
		Address uint8
	}
}

func setupDevices(config Config) (r []devices.Device) {
	for _, e := range config.Sensors {
		log.Println(e)
		dev, err := devices.Devices[e.Type](e.Address, e.Bus)
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, dev)
	}

	return r
}

func initDevices(devices []devices.Device) {
	for _, e := range devices {
		e.Init()
	}
}

func main() {
	configFile := os.Args[1]
	cfg := Config{}
	ahrs := Ahrs{}

	configData, err := ioutil.ReadFile(configFile);
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	devs := setupDevices(cfg)
	initDevices(devs)

	for {
		for _, device := range devs {
			measurement := device.Read()
			ahrs.AddMeasurement(measurement)
			fmt.Println("G-force:", ahrs.GforceTotal)
		}
	}
}
