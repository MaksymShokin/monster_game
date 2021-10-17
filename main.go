package main

import "monster_game/interactions"

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
	return ""
}

func endGame() {

}
