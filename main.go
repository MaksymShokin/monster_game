package main

import (
	"monster_game/actions"
	"monster_game/interactions"
)

var currentRound = 0
var gameRounds = []interactions.RoundData{}

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interactions.PrintGreetings()

}
func executeRound() string {
	currentRound++
	hasSpecialAttack := currentRound%3 == 0
	interactions.ShowAvailableActions(hasSpecialAttack)

	playerChoice := interactions.GetPlayerChoice(hasSpecialAttack)

	var playerAttackDamage int
	var playerHealPower int
	var monsterAttackDamage int

	if playerChoice == "ATTACK" {
		playerAttackDamage = actions.AttackMonster(false)
	} else if playerChoice == "HEAL" {
		playerHealPower = actions.HealPlayer()
	} else {
		playerAttackDamage = actions.AttackMonster(true)
	}

	monsterAttackDamage = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthValues()

	roundData := interactions.RoundData{
		Action:           playerChoice,
		PlayerAttackDmg:  playerAttackDamage,
		PlayerHeal:       playerHealPower,
		MonsterAttackDmg: monsterAttackDamage,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}

	interactions.PrintRoundData(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	}

	if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interactions.DeclareWinner(winner)
	interactions.WriteLogFile(&gameRounds)
}
