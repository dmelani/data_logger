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

const (
	dlpfCfg256Hz    byte = 0x00
	dlpfCfg188Hz    byte = 0x01
	dlpfCfg98Hz     byte = 0x02
	dlpfCfg42Hz     byte = 0x03
	dlpfCfg20Hz     byte = 0x04
	dlpfCfg10Hz     byte = 0x05
	dlpfCfg5Hz      byte = 0x06
	dlpfFsFullScale byte = 0x18
)

const (
	pwrMgmClkSelInternal        byte = 0x00
	pwrMgmClkSelXGyroRef        byte = 0x01
	pwrMgmClkSelYGyroRef        byte = 0x02
	pwrMgmClkSelZGyroRef        byte = 0x03
	pwrMgmClkSelExt32_768kHzRef byte = 0x04
	pwrMgmClkSelExt19_2kHzRef   byte = 0x05
	pwrMgmStbyZG                byte = 0x08
	pwrMgmStbyYG                byte = 0x10
	pwrMgmStbyXG                byte = 0x20
	pwrMgmSleep                 byte = 0x40
	pwrMgmHReset                byte = 0x80
)

const whoAmIMask = 0x7e

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
	if err := itg.checkWhoAmI(); err != nil {
		log.Fatal(err.Error())
	}

	itg.setRegister(regPwrMgm, pwrMgmHReset)
	itg.setRegister(regDlpfFs, dlpfFsFullScale)
	time.Sleep(50 * time.Millisecond) /* Give gyro time to settle */

	itg.setRegister(regDlpfFs, dlpfFsFullScale|dlpfCfg10Hz)
	itg.setRegister(regPwrMgm, pwrMgmClkSelXGyroRef)
}

func (itg *Itg3200) Destroy() {
}

func (itg *Itg3200) checkWhoAmI() error {
	data := []byte{0}

	itg.bus.WriteByte(regWhoAmI)
	itg.bus.Read(data)

	if data[0]&whoAmIMask != (itg.address<<1)&whoAmIMask {
		errors.New(fmt.Sprintf("ITG3200 at %x on bus %d returned wrong device id: %x\n", itg.address, itg.device, data[0]))
	}

	return nil
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

	tempC := 35 + float32(tempReg+13200)/280 // does this really make sense?
	fmt.Println("Gyro temp:", tempC)
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
