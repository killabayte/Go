package main

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}
type CpuTemp struct {
	temp []Celcius
}
type MemoryUsage struct {
	amount []Bytes
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
}
