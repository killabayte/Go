package main

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{items: make(map[P][]V), priorities: priorities}
}

func main() {

}
