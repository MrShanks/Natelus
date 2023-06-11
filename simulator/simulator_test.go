package simulator

import (
	"testing"

	"github.com/MrShanks/natelus/model"
)

func TestSimulator_DetermineFirstAttacker(t *testing.T) {
	player1 := model.NewHero()
	player2 := model.NewBeast()

	testCases := []struct {
		name         string
		player1Speed int
		player2Speed int
		player1Luck  int
		player2Luck  int
		expected     string
	}{
		{"Player1 should be the first to attack with speed higher than player 2, luck should be ignored", 10, 9, 0, 0, "Natelus"},
		{"Player2 should be the first to attack with speed higher than player 1, luck should be ignored", 9, 10, 0, 0, "Boar"},
		{"Player1 should be the first to attack with speed higher than player 2, luck should be ignored", 0, 0, 10, 9, "Natelus"},
		{"Player2 should be the first to attack with speed higher than player 1, luck should be ignored", 0, 0, 9, 10, "Boar"},
		{"When speed and luck are the same for both player player 2 should be the first to attack", 0, 0, 0, 0, "Boar"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			player1.Speed = tc.player1Speed
			player1.Luck = tc.player1Luck
			player2.Speed = tc.player2Speed
			player2.Luck = tc.player2Luck

			got, _ := DetermineFirstAttacker(&player1, &player2)

			if got.GetName() != tc.expected {
				t.Errorf("Got: %s, Expected: %s", got.GetName(), tc.expected)
			}
		})
	}
}

func TestSimulator_Swap(t *testing.T) {
	player1 := model.NewHero()
	player2 := model.NewBeast()

	testCases := []struct {
		name     string
		first    string
		second   string
		expected string
	}{
		{"Boar should be player 1", "Natelus", "Boar", "Boar"},
		{"Natelus should be player 1", "Boar", "Natelus", "Natelus"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			player1.Name = tc.first
			player2.Name = tc.second

			player1, _ := Swap(&player1, &player2)

			if player1.GetName() != tc.expected {
				t.Errorf("Got: %s, Expected: %s", player1.GetName(), tc.expected)
			}
		})
	}
}
