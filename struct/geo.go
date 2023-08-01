package geo

type Coordinates struct {
	latitude  float64
	longitude float64
}

type Landmark struct {
	name string
	Coordinates
}

func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

func (c *Coordinates) SetLongitude(longitude float64) {
	c.Longitude = longitude
}
