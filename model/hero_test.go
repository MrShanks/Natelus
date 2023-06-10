package model

import "testing"

func TestHero_MagicShield(t *testing.T) {
	hero := Hero{
		Stats{
			Name:     "Test",
			Health:   100,
			Strenght: 80,
			Defence:  60,
			Speed:    30,
			Luck:     15,
		},
	}

	damage := 10
	expectedDamage := damage / 2

	GotLucky = func(chances int) bool {
		return true
	}

	actualDamage := hero.MagicShield(damage)
	if actualDamage != expectedDamage {
		t.Errorf("MagicShield did not return the expected damage, Expected: %d, Got: %d", expectedDamage, actualDamage)
	}
}
