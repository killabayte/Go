package main

import "testing"

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
