package model

type Beast struct {
	Stats
}

func (b *Beast) Attack(defenceStat int) int {
	var damage int
	if defenceStat >= b.Strenght {
		return 0
	}

	damage = b.Strenght - defenceStat
	return damage
}

func (b *Beast) Defend(damage int) int {
	b.Health -= damage
	if b.Health <= 0 {
		b.Health = 0
	}
	return damage
}

func NewBeast() Beast {
	return Beast{*WithStats(
		Name("Boar"),
		Health(60, 90),
		Strength(60, 90),
		Defence(40, 60),
		Speed(40, 60),
		Luck(25, 40),
	)}
}
