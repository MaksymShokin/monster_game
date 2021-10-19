package interactions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHeal       int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

// func CreateRoundData(
// 	action string,
// 	playerAttackDmg int,
// 	playerHeal int,
// 	monsterAttackDmg int,
// 	playerHealth int,
// 	monsterHealth int,
// ) *RoundData {
// 	data := RoundData{
// 		Action:           action,
// 		PlayerAttackDmg:  playerAttackDmg,
// 		PlayerHeal:       playerHeal,
// 		MonsterAttackDmg: monsterAttackDmg,
// 		PlayerHealth:     playerHealth,
// 		MonsterHealth:    monsterHealth,
// 	}

// 	return &data
// }

func PrintGreetings() {
	monsterFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	monsterFigure.Print()
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
	monsterFigure := figure.NewColorFigure("GAME OVER!", "", "green", true)
	monsterFigure.Print()
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

func WriteLogFile(rounds *[]RoundData) {
	path, err := os.Executable()

	if err != nil {
		fmt.Println("Getting path failed")
		return
	}

	path = filepath.Dir(path)

	file, err := os.Create(path + "/gamelog.txt")

	if err != nil {
		fmt.Println("Writing log failed")
		return
	}

	for index, value := range *rounds {
		round := fmt.Sprintf(
			"Round: %v\nAction: %v\nPlayer Attack Dmg: %v\nMonster Attack Dmg %v\nPlayer Heal %v\nPlayer Health %v\nMonster Health %v\n-----------------------\n",
			index+1,
			value.Action,
			value.PlayerAttackDmg,
			value.MonsterAttackDmg,
			value.PlayerHeal,
			value.PlayerHealth,
			value.MonsterHealth,
		)

		_, err := file.WriteString(round)

		if err != nil {
			fmt.Println("Writing to file failed")
			continue
		}
	}

	file.Close()
}
