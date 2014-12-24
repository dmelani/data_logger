package devices

import (
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

const deviceID byte = 0xE5

type Adxl345 struct {
	bus     *i2c.I2C
	device  int
	address uint8
}

func NewAdxl345(address uint8, device int) (*Adxl345, error) {
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
	data := []byte{0}

	adxl.bus.WriteByte(regDevid)
	adxl.bus.Read(data)
	log.Println(data)

	if data[0] != deviceID {
		log.Fatalf("ADXL345 at %x on bus %d returned wrong device id: %x\n", adxl.address, adxl.device, data[0])
	}
}
