package model

import (
	"fmt"
)

type Hero struct {
	Stats
}

func NewHero() Hero {
	return Hero{*WithStats(
		Name("Natelus"),
		Health(70, 100),
		Strength(79, 80),
		Defence(45, 55),
		Speed(40, 50),
		Luck(10, 30),
	)}
}

func (h *Hero) Attack(defenceStat int) int {
	if defenceStat >= h.Strenght {
		return 0
	}
	damage := h.Strenght - defenceStat
	damage = h.RapidStrike(damage)
	return damage
}

func (h *Hero) Defend(damage int) int {
	actualDamage := h.MagicShield(damage)
	h.Health -= actualDamage
	if h.Health <= 0 {
		h.Health = 0
	}
	return actualDamage
}

func (c *Hero) RapidStrike(damage int) int {
	// 10% chance
	if GotLucky(10) {
		fmt.Printf("%s uses RapidStrike, attacking twice.\n", c.GetName())
		return damage * 2
	}
	return damage
}

func (c *Hero) MagicShield(damage int) int {
	// 20% chance
	if GotLucky(20) {
		fmt.Printf("Inflicting %d damages\n", damage)
		fmt.Printf("%s uses MagicShield, only taking half of the damage\n", c.GetName())
		return damage / 2
	}
	return damage
}
