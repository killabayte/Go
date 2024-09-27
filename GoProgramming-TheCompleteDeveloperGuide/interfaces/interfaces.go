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

func main() {

}
