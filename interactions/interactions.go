package interactions

import "fmt"

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHeal       int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreetings() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("moster is comming...")
	fmt.Println("GOOD LUCK")
}

func ShowAvailableActions(hasSpecialAttack bool) {
	fmt.Println("Make you decision:")
	fmt.Println("------------------")
	fmt.Println("1. Attack monster")
	fmt.Println("2. Heal")

	if hasSpecialAttack {
		fmt.Println("3. Perform special attack")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("------------------")
	fmt.Printf("%v WON THE BATTLE \n", winner)
}

func PrintRoundData(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v. \n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "HEAL" {
		fmt.Printf("Player healed for %v. \n", roundData.PlayerHeal)
	} else {
		fmt.Printf("Player performed strong attack against monster for %v. \n", roundData.PlayerAttackDmg)

	}

	fmt.Printf("Monster attacked player for %v. \n", roundData.MonsterAttackDmg)
	fmt.Printf("Player health: %v. \n", roundData.PlayerHealth)
	fmt.Printf("Monster health: %v. \n", roundData.MonsterHealth)
}
