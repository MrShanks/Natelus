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

func TestHero_Attack(t *testing.T) {
	hero := Hero{
		Stats{
			Name:     "Natelus",
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
		{"Attack with 70 Strenght to a defender with 60 defence and luck should deal 20 damage", 60, 20, true},
		{"Attack with 70 Strenght to a defender with 70 defence and luck should deal 0 damage", 70, 0, true},
		{"Attack with 70 Strenght to a defender with 60 defence and no luck should deal 10 damage", 60, 10, false},
		{"Attack with 70 Strenght to a defender with 70 defence and no luck should deal 0 damage", 70, 0, false},
		{"Attack with less Strenght than defender defence should deal 0 damage", 80, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GotLucky = func(chances int) bool {
				return tc.luck
			}
			actualDamage := hero.Attack(tc.defenceStat)
			if actualDamage != tc.expectedDamage {
				t.Errorf("Got: %d, Expected: %d", actualDamage, tc.expectedDamage)
			}
		})
	}
}

func TestHero_Defend(t *testing.T) {
	hero := Hero{
		Stats{
			Name:     "Natelus",
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
		{"Defend an Attack of 10 damage should deal 5 damage with luck", 10, 5, true},
		{"Defend an Attack of 0 damage should deal 0 damage with luck", 0, 0, true},
		{"Defend an Attack of 10 damage should deal 10 damage with no luck", 10, 10, false},
		{"Defend an Attack of 0 damage should deal 0 damage with no luck", 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GotLucky = func(chances int) bool {
				return tc.luck
			}
			actualDamage := hero.Defend(tc.damage)
			if actualDamage != tc.expectedDamage {
				t.Errorf("Got: %d, Expected: %d", actualDamage, tc.expectedDamage)
			}
		})
	}
}
