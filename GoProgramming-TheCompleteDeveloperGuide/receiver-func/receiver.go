package main

import "fmt"

type Player struct {
	name      string
	MaxHealth int
	MaxEnergy int
	Health    int
	Energy    int
}

func NewPlayer() Player {
	return Player{
		MaxHealth: 5,
		MaxEnergy: 10,
	}
}

func (p *Player) modifyHealth(modifyValueLow int, modifyValueTop int) {
	if modifyValueLow != 0 {
		p.MaxHealth = p.MaxHealth - modifyValueLow
	} else if modifyValueTop != 0 {
		p.MaxHealth = p.MaxHealth + modifyValueTop
	}
}

func main() {
	erik := NewPlayer()
	fmt.Println("Player's maxHealth: ", erik.MaxHealth)
	fmt.Println("Player's maxEnergy: ", erik.MaxEnergy)
}
