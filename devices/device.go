package devices

type Device interface {
	Init()
	Destroy()
	Read() Measurement
}

var Devices = map[string]func(uint8, int) (Device, error){
	"adxl345":  NewAdxl345,
	"itg3200":  NewItg3200,
	"hmc5883l": NewHmc5883l,
}
