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
		p.Health = p.MaxHealth - modifyValueLow
		if p.Health < 0 {
			p.Health = 0
		}
	} else if modifyValueTop != 0 {
		p.Health = p.MaxHealth + modifyValueTop
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
	}
}

func (p *Player) modifyEnergy(modifyValueLow int, modifyValueTop int) {
	if modifyValueLow != 0 {
		p.Energy = p.MaxEnergy - modifyValueLow
		if p.Energy < 0 {
			p.Energy = 0
		}
	} else if modifyValueTop != 0 {
		p.Energy = p.MaxEnergy + modifyValueTop
		if p.Energy > p.MaxEnergy {
			p.Energy = p.MaxEnergy
		}
	}
}

func main() {
	erik := NewPlayer()
	fmt.Println("Player's maxHealth before start: ", erik.MaxHealth)
	fmt.Println("Player's maxEnergy before start: ", erik.MaxEnergy)

	erik.modifyHealth(3, 0)
	erik.modifyEnergy(7, 0)
	fmt.Println("Player's maxHealth, after first boss: ", erik.Health)
	fmt.Println("Player's maxEnergy, after first boos: ", erik.Energy)

	//Drick some bottle
	erik.modifyHealth(0, 8)
	erik.modifyEnergy(0, 5)
	fmt.Println("Player's maxHealth, after bottle: ", erik.Health)
	fmt.Println("Player's maxEnergy, after bottle: ", erik.Energy)
}
