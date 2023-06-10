package model

type Attacker interface {
	Attack(int) int
}

type Defender interface {
	Defend(int) int
}

type AttackerDefender interface {
	Attacker
	Defender
}

type Character interface {
	AttackerDefender
	GetName() string
	GetDefence() int
	GetHealth() int
	GetLuck() int
	GetSpeed() int
}
