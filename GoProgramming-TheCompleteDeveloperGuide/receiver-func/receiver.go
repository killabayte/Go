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
func (p *Player) modifyEnergy(modifyValueLow int, modifyValueTop int) {
	if modifyValueLow != 0 {
		p.MaxEnergy = p.MaxEnergy - modifyValueLow
	} else if modifyValueTop != 0 {
		p.MaxEnergy = p.MaxEnergy + modifyValueTop
	}
}

func main() {
	erik := NewPlayer()
	fmt.Println("Player's maxHealth before start: ", erik.MaxHealth)
	fmt.Println("Player's maxEnergy before start: ", erik.MaxEnergy)

	erik.modifyHealth(3, 0)
	erik.modifyEnergy(7, 0)
	fmt.Println("Player's maxHealth, after first boss: ", erik.MaxHealth)
	fmt.Println("Player's maxEnergy, after first boos: ", erik.MaxEnergy)

	//Drick some bottle
	erik.modifyHealth(0, 2)
	erik.modifyEnergy(0, 5)
	fmt.Println("Player's maxHealth, after bottle: ", erik.MaxHealth)
	fmt.Println("Player's maxEnergy, after bottle: ", erik.MaxEnergy)

}
