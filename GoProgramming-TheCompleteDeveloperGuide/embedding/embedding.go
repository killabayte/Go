package main

import "fmt"

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

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}
func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(c.temp); i++ {
		sum += int(c.temp[i])
	}
	return sum / len(c.temp)
}
func (m *MemoryUsage) AverageMemUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageBandwidth())
	fmt.Printf("Average temp: %v\n", dash.AverageCpuTemp())
	fmt.Printf("Avarage memory usage: %v\n", dash.AverageMemUsage())
}
