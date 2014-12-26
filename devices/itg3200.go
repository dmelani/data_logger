package devices

import (
	"bytes"
	"encoding/binary"
//	"errors"
	"fmt"
	i2c "github.com/davecheney/i2c"
)

const (
	regWhoAmI    = 0x00
	regSmplrtDiv = 0x15
	regDlpfFs    = 0x16
	regIntCfg    = 0x17
	regIntStatus = 0x1a
	regTempOutH  = 0x1b
	regTempOutL  = 0x1c
	regGyroXoutH = 0x1d
	regGyroXoutL = 0x1e
	regGyroYoutH = 0x1f
	regGyroYoutL = 0x20
	regGyroZoutH = 0x21
	regGyroZoutL = 0x22
	regPwrMgm    = 0x3e
)

type Itg3200 struct {
	bus     *i2c.I2C
	device  int
	address uint8
}

func NewItg3200(address uint8, device int) (Device, error) {
	itg := Itg3200{
		device:  device,
		address: address,
	}

	bus, err := i2c.New(address, device)
	if err != nil {
		return nil, err
	}

	itg.bus = bus
	return &itg, nil
}

func (itg *Itg3200) Init() {
}

func (itg *Itg3200) Destroy() {
}

func (itg *Itg3200) Read() Measurement {
	data := make([]byte, 8, 8)
	var tempReg int16
	var xReg int16
	var yReg int16
	var zReg int16

	itg.bus.WriteByte(regTempOutH)
	itg.bus.Read(data)

	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &tempReg)
	binary.Read(buf, binary.BigEndian, &xReg)
	binary.Read(buf, binary.BigEndian, &yReg)
	binary.Read(buf, binary.BigEndian, &zReg)

	fmt.Println("Temp:", tempReg)
	ret := &Gyro{}
	ret.data[0] = float32(xReg)
	ret.data[1] = float32(yReg)
	ret.data[2] = float32(zReg)

	return ret
}

func (itg *Itg3200) setRegister(register byte, flags byte) {
	data := []byte{register, flags}

	itg.bus.Write(data)
}
