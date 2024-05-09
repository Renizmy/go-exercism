package elon

import "fmt"

// Drive the car
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayDistance return the total distance
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery return the current battery state
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish calculate if the car can finish the track
func (car *Car) CanFinish(trackDistance int) bool {
	return car.battery/car.batteryDrain*car.speed >= trackDistance
}
