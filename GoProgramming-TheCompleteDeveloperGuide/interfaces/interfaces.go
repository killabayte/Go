package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int
type LiftPicker interface {
	PickLift() Lift
}

type Motocycle string
type Car string
type Truck string

func (m Motocycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motocycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("Send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("Send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("Send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motocycle := Motocycle("Croozer")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motocycle)
}
