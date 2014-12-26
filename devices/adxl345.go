package devices

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	i2c "github.com/davecheney/i2c"
	"log"
)

const (
	regDevid         = 0x00
	regThreshTap     = 0x1d
	regOfsX          = 0x1e
	regOfsY          = 0x1f
	regOfsZ          = 0x20
	regDur           = 0x21
	regLatent        = 0x22
	regWindow        = 0x23
	regThreshAct     = 0x24
	regThreshInact   = 0x25
	regTimeInact     = 0x26
	regActInact_Ctl  = 0x27
	regThreshFF      = 0x28
	regTimeFF        = 0x29
	regTapAxes       = 0x2a
	regActTap_Status = 0x2b
	regBWRate        = 0x2c
	regPowerCtl      = 0x2d
	regIntEnable     = 0x2e
	regIntMap        = 0x2f
	regIntSource     = 0x30
	regDataFormat    = 0x31
	regDataX0        = 0x32
	regDataX1        = 0x33
	regDataY0        = 0x34
	regDataY1        = 0x35
	regDataZ0        = 0x36
	regDataZ1        = 0x37
	regFifoCtl       = 0x38
	regFifoStatus    = 0x39
)

const (
	powerCtl8Hz       byte = 0x00
	powerCtl4Hz       byte = 0x01
	powerCtl2Hz       byte = 0x02
	powerCtl1Hz       byte = 0x03
	powerCtlSleep     byte = 0x04
	powerCtlMeasure   byte = 0x08
	powerCtlAutoSleep byte = 0x10
	powerCtlLink      byte = 0x20
)

const (
	dataFormatRange2g   byte = 0x00
	dataFormatRange4g   byte = 0x01
	dataFormatRange8g   byte = 0x02
	dataFormatRange16g  byte = 0x03
	dataFormatJustify   byte = 0x04
	dataFormatFullRes   byte = 0x08
	dataFormatIntInvert byte = 0x20
	dataFormatSpi       byte = 0x40
	dataFormatSelfTest  byte = 0x80
)

const (
	bwRate6_25 byte = 0x06
	bwRate12_5 byte = 0x07
	bwRate25   byte = 0x08
	bwRate50   byte = 0x09
	bwRate100  byte = 0x0a
	bwRate200  byte = 0x0b
	bwRate400  byte = 0x0c
	bwRate800  byte = 0x0d
	bwRate1600 byte = 0x0e
	bwRate3200 byte = 0x0f
)

const deviceID byte = 0xE5

const fullResolutionScaleFactor float32 = 3.9

type Adxl345 struct {
	bus     *i2c.I2C
	device  int
	address uint8
}

func NewAdxl345(address uint8, device int) (Device, error) {
	adxl := Adxl345{
		device:  device,
		address: address,
	}

	bus, err := i2c.New(address, device)
	if err != nil {
		return nil, err
	}

	adxl.bus = bus
	log.Println(adxl.bus)
	return &adxl, nil
}

func (adxl *Adxl345) Init() {
	if err := adxl.checkDevID(); err != nil {
		log.Fatalf(err.Error())
	}

	adxl.setRegister(regDataFormat, dataFormatRange16g|dataFormatFullRes)
	adxl.setRegister(regBWRate, bwRate400)
	adxl.setRegister(regPowerCtl, powerCtlMeasure)
}

func (adxl *Adxl345) Destroy() {
}

func (adxl *Adxl345) Read() Measurement {
	data := make([]byte, 6, 6)
	var xReg int16
	var yReg int16
	var zReg int16

	adxl.bus.WriteByte(regDataX0)
	adxl.bus.Read(data)

	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &xReg)
	binary.Read(buf, binary.LittleEndian, &yReg)
	binary.Read(buf, binary.LittleEndian, &zReg)

	ret := &Acceleration{}
	ret.data[0] = int32(float32(xReg) * fullResolutionScaleFactor)
	ret.data[1] = int32(float32(yReg) * fullResolutionScaleFactor)
	ret.data[2] = int32(float32(zReg) * fullResolutionScaleFactor)

	return ret
}

func (adxl *Adxl345) checkDevID() error {
	data := []byte{0}

	adxl.bus.WriteByte(regDevid)
	adxl.bus.Read(data)

	if data[0] != deviceID {
		errors.New(fmt.Sprintf("ADXL345 at %x on bus %d returned wrong device id: %x\n", adxl.address, adxl.device, data[0]))
	}

	return nil
}

func (adxl *Adxl345) setRegister(register byte, flags byte) {
	data := []byte{register, flags}

	adxl.bus.Write(data)
}
