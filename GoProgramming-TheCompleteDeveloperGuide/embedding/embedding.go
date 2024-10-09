package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}
func (s Shipping) String() string {
	return []string{"Groud", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}
type Shipper interface {
	Ship() Shipper
}
type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}
func (s *SpamMail) Ship() Shipping {
	return Air
}
func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxixWaste struct {
	weight int
}

func main() {

}
