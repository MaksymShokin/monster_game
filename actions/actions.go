package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var monsterHealth = MONSTER_STARTING_HEALTH
var playerHealth = PLAYER_STARTING_HEALTH

func AttackMonster(isSpecial bool)  {
	var attackPower int
	if isSpecial {
		attackPower = generateRandBetween(10, 25)
	} else {
		attackPower = generateRandBetween(15, 35)
	}

	monsterHealth -= attackPower

}

func HealPlayer() {
	healPower := generateRandBetween(20, 50)

	playerHealth += healPower

	if playerHealth > PLAYER_STARTING_HEALTH {
		playerHealth = PLAYER_STARTING_HEALTH
	}

}

func AttackPlayer() (int, int) {
	attackPower := generateRandBetween(15, 30)

	playerHealth -= attackPower

	return playerHealth, monsterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
