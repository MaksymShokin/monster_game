package interactions

import "fmt"

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
