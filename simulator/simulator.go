package simulator

import (
	"fmt"
	"time"

	"github.com/MrShanks/natelus/model"
)

const Rounds int = 20

func Swap(attacker, defender model.Character) (model.Character, model.Character) {
	return defender, attacker
}

func DetermineFirstAttacker(player1, player2 model.Character) (model.Character, model.Character) {
	if player1.GetSpeed() > player2.GetSpeed() || (player1.GetSpeed() == player2.GetSpeed() && player1.GetLuck() > player2.GetLuck()) {
		return player1, player2
	}
	return player2, player1
}

func Battle(attacker, defender model.Character) {
	fmt.Printf("Let the battle begin\n\n")

	attacker, defender = DetermineFirstAttacker(attacker, defender)

	for i := 1; i <= Rounds; i++ {
		time.Sleep(2000 * time.Millisecond)
		// Check if someone died
		if attacker.GetHealth() <= 0 || defender.GetHealth() <= 0 {
			break
		}

		fmt.Printf("Round %d: ", i)

		// Calculate the attack damage attacker strength - defencer defence
		fmt.Printf("%s Attacks\n", attacker.GetName())
		damage := attacker.Attack(defender.GetDefence())

		// Roll the dice to determine if the defender is lucky and dodges the attack
		if model.GotLucky(defender.GetLuck()) {
			damage = 0
			fmt.Printf("%s got lucky, they dodged the attack ", defender.GetName())
		}

		// Defender can have special abilities to limit the damages
		actualDamage := defender.Defend(damage)

		fmt.Printf("%d damage were dealt\n", actualDamage)
		fmt.Printf("%s, has now: %d HP\n\n", defender.GetName(), defender.GetHealth())

		attacker, defender = Swap(attacker, defender)
	}

}
