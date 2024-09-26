package main

import (
	"testing"
)

func TestMaxResources(t *testing.T) {
	np := NewPlayer()

	defaultMaxHealth := np.MaxHealth
	defaultMaxEnergy := np.MaxEnergy

	np.modifyHealth(0, np.MaxHealth+1)
	np.modifyEnergy(0, np.MaxEnergy+1)

	if np.MaxHealth > defaultMaxHealth {
		t.Errorf("Max health: %v is bigger than maximum %v", np.MaxHealth, defaultMaxHealth)
	}
	if np.MaxEnergy > defaultMaxEnergy {
		t.Errorf("Max energy: %v is bigger than maximum %v", np.MaxEnergy, defaultMaxEnergy)
	}
}

func TestMinResources(t *testing.T) {
	np := NewPlayer()

	np.modifyHealth(np.MaxHealth+1, 0)
	np.modifyEnergy(np.MaxEnergy+1, 0)

	if np.Health < 0 {
		t.Errorf("Players health: %v is lower than zero", np.Health)
	}
	if np.Energy < 0 {
		t.Errorf("Players energy: %v is lower than zero", np.Energy)
	}
}
