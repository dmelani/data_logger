package devices

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	i2c "github.com/davecheney/i2c"
	"log"
	"time"
)

const (
	regConfigA = 0x00
	regConfigB = 0x01
	regMode    = 0x02
	regXMsb    = 0x03
	regXLsb    = 0x04
	regZMsb    = 0x05
	regZLsb    = 0x06
	regYMsb    = 0x07
	regYLsb    = 0x08
	regStatus  = 0x09
	regIdentA  = 0x0a
	regIdentB  = 0x0b
	regIdentC  = 0x0c
)

const (
	identA byte = 0x48
	identB byte = 0x34
	identC byte = 0x33
)

const (
	configAMeasAvg1     byte = 0x00 << 5
	configAMeasAvg2     byte = 0x01 << 5
	configAMeasAvg4     byte = 0x02 << 5
	configAMeasAvg8     byte = 0x03 << 5
	configAOutRate0_75  byte = 0x00 << 2
	configAOutRate1_5   byte = 0x01 << 2
	configAOutRate3     byte = 0x02 << 2
	configAOutRate7_5   byte = 0x03 << 2
	configAOutRate15    byte = 0x04 << 2
	configAOutRate30    byte = 0x05 << 2
	configAOutRate75    byte = 0x06 << 2
	configAMeasNormal   byte = 0x00
	configAMeasPositive byte = 0x01
	configAMeasNegative byte = 0x02
)

const (
	configBGain0 byte = 0x00 << 5
	configBGain1 byte = 0x01 << 5
	configBGain2 byte = 0x02 << 5
	configBGain3 byte = 0x03 << 5
	configBGain4 byte = 0x04 << 5
	configBGain5 byte = 0x05 << 5
	configBGain6 byte = 0x06 << 5
	configBGain7 byte = 0x07 << 5
)

const (
	modeContinuous byte = 0x00
	modeSingle     byte = 0x01
	modeIdle       byte = 0x02
)

var scale = map[byte]float32{
	configBGain0: 0.73,
	configBGain1: 0.92,
	configBGain2: 1.22,
	configBGain3: 1.52,
	configBGain4: 2.27,
	configBGain5: 2.56,
	configBGain6: 3.03,
	configBGain7: 4.35,
}

type Hmc5883l struct {
	bus     *i2c.I2C
	device  int
	address uint8
	gain    byte
}

func NewHmc5883l(address uint8, device int) (Device, error) {
	hmc := Hmc5883l{
		device:  device,
		address: address,
	}

	bus, err := i2c.New(address, device)
	if err != nil {
		return nil, err
	}

	hmc.bus = bus
	return &hmc, nil
}

func (hmc *Hmc5883l) Init() {
	if err := hmc.checkIdent(); err != nil {
		log.Fatal(err.Error())
	}

	hmc.gain = configBGain2
	hmc.setRegister(regConfigA, configAMeasAvg8|configAOutRate75|configAMeasNormal)
	hmc.setRegister(regConfigB, hmc.gain)
	hmc.setRegister(regMode, modeContinuous)

	time.Sleep(6 * time.Millisecond)
}

func (hmc *Hmc5883l) Destroy() {
}

func (hmc *Hmc5883l) checkIdent() error {
	data := make([]byte, 3, 3)

	hmc.bus.WriteByte(regWhoAmI)
	hmc.bus.Read(data)

	if data[0] != identA || data[1] != identB || data[2] != identC {
		errors.New(fmt.Sprintf("hmc5883l at %x on bus %d returned wrong identity %x %x %x:\n", hmc.address, hmc.device, data[0], data[1], data[2]))
	}

	return nil
}

func (hmc *Hmc5883l) Read() Measurement {
	data := make([]byte, 6, 6)
	var xReg int16
	var yReg int16
	var zReg int16

	hmc.bus.WriteByte(regXMsb)
	hmc.bus.Read(data)

	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &xReg)
	binary.Read(buf, binary.BigEndian, &zReg)
	binary.Read(buf, binary.BigEndian, &yReg)

	ret := &MagneticField{}
	ret.data[0] = int32(float32(xReg) * scale[hmc.gain])
	ret.data[1] = int32(float32(yReg) * scale[hmc.gain])
	ret.data[2] = int32(float32(zReg) * scale[hmc.gain])

	return ret
}

func (hmc *Hmc5883l) setRegister(register byte, flags byte) {
	data := []byte{register, flags}

	hmc.bus.Write(data)
}
