package main

import (
	"monster_game/actions"
	"monster_game/interactions"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interactions.PrintGreetings()

}
func executeRound() string {
	currentRound++
	hasSpecialAttack := currentRound%3 == 0
	interactions.ShowAvailableActions(hasSpecialAttack)

	playerChoice := interactions.GetPlayerChoice(hasSpecialAttack)

	var playerHealth int
	var monsterHealth int

	if playerChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if playerChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}

	playerHealth, monsterHealth = actions.AttackPlayer()

	if playerHealth <= 0 {
		return "Monster"
	}

	if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {

}
