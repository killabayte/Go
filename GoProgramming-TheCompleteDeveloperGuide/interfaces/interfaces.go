package main

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

func main() {

}
