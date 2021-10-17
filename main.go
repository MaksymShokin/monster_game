package main

import "monster_game/interactions"

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

	
	return ""
}

func endGame() {

}