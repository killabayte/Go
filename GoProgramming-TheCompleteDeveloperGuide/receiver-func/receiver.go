package main

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}
