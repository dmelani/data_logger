Data Logger
===========
This is going to a data logger for track days and trips. It is currently based on a Beaglebone Black with a SparkFun 9 Degrees of Freedom sensor stick connected to it.

####Supported hardware
The data logger currently has support for the following i2c sensors

#####Gyroscopes
- itg3200 

#####Magnetic field sensors
- hmc5883l

#####Accelerometers
- adxl345

####Current status
Only the basic functionalities of the sensors are in place, and there are many features that still have to be implemented.

I still have to implement at least the following:
- Calibration of all sensors
- Temperature compensation for itg32000
- An entity that collects data an calculates current heading and lean angle
- A logging framework that collects and logs data

Look at this document for hints on how calibration of hmc5883l should be done http://www.joics.com/publishedpapers/2013_10_6_1551_1558.pdf

For error correction in itg3200, look at this http://mathworld.wolfram.com/LeastSquaresFitting.html and http://en.wikipedia.org/wiki/Errors-in-variables_models

Might want to try out a PID controller too even it if is not direclty relevant to this project at the moment: https://github.com/felixge/pidctrl and http://en.wikipedia.org/wiki/PID_controller

For info on developing the IMU part: http://www.starlino.com/imu_guide.html
This might also help: http://stackoverflow.com/questions/1586658/combine-gyroscope-and-accelerometer-data
