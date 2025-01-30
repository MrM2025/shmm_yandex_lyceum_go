package main

import "fmt"

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	GetType() string
}

type Car struct {
	Speed    float64
	Type     string
	FuelType string
}

type Motorcycle struct {
	Speed float64
	Type  string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

func (c Car) GetType() string {
	return c.Type
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func (m Motorcycle) GetType() string {
	return m.Type
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {

	travelTimes := make(map[string]float64)

	for _, vehicle := range vehicles {
		travelTimes[vehicle.GetType()] = vehicle.CalculateTravelTime(distance)
	}

	return travelTimes
}

func main() {

	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}
	vehicles := []Vehicle{car, motorcycle}
	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)

	fmt.Println(travelTimes)
}
