package main

import (
	i2c "github.com/davecheney/i2c"
	"log"
)

type Adxl345 struct {
	bus *i2c.I2C
	device int
	address uint8
}

func NewAdxl345(address uint8, device int) (*Adxl345, error) {
	adxl := Adxl345{
		device : device,
		address : address,
	}

	bus, err := i2c.New(address, device)
	if err != nil {
		return nil, err
	}

	adxl.bus = bus
	log.Println(adxl.bus)
	return &adxl, nil
}

func (adxl Adxl345) Init() {
}
