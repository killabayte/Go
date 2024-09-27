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

func main() {

}
