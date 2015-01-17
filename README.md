Data Logger
===========

##Overview
This is going to a data logger for track days and trips. It is currently based on a Beaglebone Black with a SparkFun 9 Degrees of Freedom sensor stick connected to it.

##Supported hardware
The data logger currently has support for the following i2c sensors

###Gyroscopes
- itg3200 

###Magnetic field sensors
- hmc5883l

###Accelerometers
- adxl345

##Current status
Only the basic functionalities of the sensors are in place, and there are many features that still have to be implemented.

I still have to implement at least the following:
- Calibration of all sensors
- Temperature compensation for itg32000
- An entity that collects data an calculates current heading and lean angle
- A logging framework that collects and logs data
