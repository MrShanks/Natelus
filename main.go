package main

import (
	"github.com/MrShanks/natelus/model"
	"github.com/MrShanks/natelus/simulator"
)

func main() {
	natelus := model.NewHero()
	boar := model.NewBeast()

	simulator.Battle(&natelus, &boar)
}
