package model

import "testing"

func TestHero_MagicShield(t *testing.T) {
	hero := NewHero()

	testCases := []struct {
		name           string
		damage         int
		expectedDamage int
		luck           bool
	}{
		{"Damage 10 with MagicShield ability should only deal 5 damage", 10, 5, true},
		{"Damage 1 with MagicShield ability should only deal 0 damage", 1, 0, true},
		{"Damage 0 with MagicShield ability should deal 0 damage", 0, 0, true},
		{"Damage 10 without MagicShield ability should deal 10 damage", 10, 10, false},
		{"Damage 1 without MagicShield ability should only deal 1 damage", 1, 1, false},
		{"Damage 0 without MagicShield ability should deal 0 damage", 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GotLucky = func(chances int) bool {
				return tc.luck
			}
			actualDamage := hero.MagicShield(tc.damage)
			if actualDamage != tc.expectedDamage {
				t.Errorf("Got: %d, Expected: %d", actualDamage, tc.expectedDamage)
			}
		})
	}
}

func TestHero_RapidStrike(t *testing.T) {
	hero := NewHero()

	testCases := []struct {
		name           string
		damage         int
		expectedDamage int
		luck           bool
	}{
		{"Damage 10 with RapidStrike ability should deal 20 damage", 10, 20, true},
		{"Damage 0 with RapidStrike ability should deal 0 damage", 0, 0, true},
		{"Damage 10 without RapidStrike ability should deal 10 damage", 10, 10, false},
		{"Damage 0 without RapidStrike ability should deal 0 damage", 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GotLucky = func(chances int) bool {
				return tc.luck
			}
			actualDamage := hero.RapidStrike(tc.damage)
			if actualDamage != tc.expectedDamage {
				t.Errorf("Got: %d, Expected: %d", actualDamage, tc.expectedDamage)
			}
		})
	}
}
