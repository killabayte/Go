package main

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

func main() {

}
