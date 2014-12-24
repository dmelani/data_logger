package main

import (
	//i2c "github.com/davecheney/i2c"
	"log"
)

func main() {
	/*
	itg, err := i2c.New(0x68, 1)
	if err != nil { log.Fatal(err) }
	adxl, err := i2c.New(0x53, 1)
	if err != nil { log.Fatal(err) }
	hmc, err := i2c.New(0x1e, 1)
	if err != nil { log.Fatal(err) }
	log.Print(itg)
	log.Print(hmc)
	*/
	adxl, err := NewAdxl345(0x53, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(adxl)
}

