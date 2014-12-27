package devices

type Measurement interface {
	Value() interface{}
}

/*
 * Acceleration
 */
type Acceleration struct {
	data [3]int32 /* mg */
}

func (a *Acceleration) Value() interface{} {
	return a.data
}

/*
 * Magnetic Field
 */
type MagneticField struct {
	data [3]float64 /* milligauss */
}

func (mf *MagneticField) Value() interface{} {
	return mf.data
}

/*
 * Gyroscope
 */
type Gyro struct {
	data [3]int32 /* millidegrees/s */
}

func (g *Gyro) Value() interface{} {
	return g.data
}
