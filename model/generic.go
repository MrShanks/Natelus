package model

import (
	"math/rand"
	"time"
)

type Stats struct {
	Name     string
	Health   int
	Strenght int
	Defence  int
	Speed    int
	Luck     int
}

func WithStats(options ...func(*Stats)) *Stats {
	stats := new(Stats)
	for _, opt := range options {
		opt(stats)
	}
	return stats
}

func Name(name string) func(*Stats) {
	return func(stats *Stats) {
		stats.Name = name
	}
}

func Health(min, max int) func(*Stats) {
	rand.Seed(time.Now().UnixNano())
	return func(stats *Stats) {
		stats.Health = rand.Intn(max-min) + min
	}
}

func Strength(min, max int) func(*Stats) {
	rand.Seed(time.Now().UnixNano())
	return func(stats *Stats) {
		stats.Strenght = rand.Intn(max-min) + min
	}
}

func Defence(min, max int) func(*Stats) {
	rand.Seed(time.Now().UnixNano())
	return func(stats *Stats) {
		stats.Defence = rand.Intn(max-min) + min
	}
}

func Speed(min, max int) func(*Stats) {
	rand.Seed(time.Now().UnixNano())
	return func(stats *Stats) {
		stats.Speed = rand.Intn(max-min) + min
	}
}

func Luck(min, max int) func(*Stats) {
	rand.Seed(time.Now().UnixNano())
	return func(stats *Stats) {
		stats.Luck = rand.Intn(max-min) + min
	}
}

func (s *Stats) GetName() string {
	return s.Name
}

func (s *Stats) GetDefence() int {
	return s.Defence
}

func (s *Stats) GetHealth() int {
	return s.Health
}

func (s *Stats) GetLuck() int {
	return s.Luck
}

func (s *Stats) GetSpeed() int {
	return s.Speed
}

type LuckyFunc func(chances int) bool

var GotLucky LuckyFunc = func(chances int) bool {
	rand.Seed(time.Now().UnixNano())
	return chances >= rand.Intn(100)
}
