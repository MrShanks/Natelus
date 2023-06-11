package model

import "testing"

func TestBeast_Attack(t *testing.T) {
	beast := Beast{
		Stats{
			Name:     "Boar",
			Strenght: 70,
			Health:   90,
			Defence:  55,
			Speed:    40,
			Luck:     15,
		},
	}

	testCases := []struct {
		name           string
		defenceStat    int
		expectedDamage int
		luck           bool
	}{
		{"Attack with 70 Strenght to a defender with 60 defence should deal 10 damage", 60, 10, false},
		{"Attack with 70 Strenght to a defender with 70 defence should deal 0 damage", 70, 0, false},
		{"Attack with less Strenght than defender defence should deal 0 damage", 80, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GotLucky = func(chances int) bool {
				return tc.luck
			}
			actualDamage := beast.Attack(tc.defenceStat)
			if actualDamage != tc.expectedDamage {
				t.Errorf("Got: %d, Expected: %d", actualDamage, tc.expectedDamage)
			}
		})
	}
}

func TestBeast_Defend(t *testing.T) {
	beast := Beast{
		Stats{
			Name:     "Boar",
			Strenght: 70,
			Health:   90,
			Defence:  55,
			Speed:    40,
			Luck:     15,
		},
	}

	testCases := []struct {
		name           string
		damage         int
		expectedDamage int
		luck           bool
	}{
		{"Defend an Attack of 10 damage should deal 10 damage", 10, 10, false},
		{"Defend an Attack of 0 damage should deal 0 damage", 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GotLucky = func(chances int) bool {
				return tc.luck
			}
			actualDamage := beast.Defend(tc.damage)
			if actualDamage != tc.expectedDamage {
				t.Errorf("Got: %d, Expected: %d", actualDamage, tc.expectedDamage)
			}
		})
	}
}
