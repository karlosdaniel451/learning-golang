package main

import "fmt"

const (
	InitialBattery = 0
)

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      InitialBattery,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// Updates the number of meters driven based on the car's speed, and reduces the
// battery according to the battery drainage.
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}

	car.distance -= car.speed
	car.battery -= car.batteryDrain
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// Check if a car can finish a race based on its properties and the race's track distance.
func (car *Car) CanFinish(trackDistance int) bool {
	return car.Range() >= trackDistance
}

// Get how far the car can go on a single charge.
func (car *Car) Range() int {
	return InitialBattery / car.batteryDrain * car.speed
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Printf("%#v\n", car)

	distance := 800
	track := NewTrack(distance)
	fmt.Printf("%#v\n", track)

	speed = 5
	batteryDrain = 2
	car = NewCar(speed, batteryDrain)

	trackDistance := 100

	fmt.Println(car.CanFinish(trackDistance))
}
