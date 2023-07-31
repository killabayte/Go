package main

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {

}
