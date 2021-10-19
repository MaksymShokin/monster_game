package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var monsterHealth = MONSTER_STARTING_HEALTH
var playerHealth = PLAYER_STARTING_HEALTH

func AttackMonster(isSpecial bool) int {
	var attackPower int
	if isSpecial {
		attackPower = generateRandBetween(10, 25)
	} else {
		attackPower = generateRandBetween(15, 35)
	}

	monsterHealth -= attackPower

	return attackPower
}

func HealPlayer() int {
	healPower := generateRandBetween(15, 55)

	if playerHealth+healPower > 200 {
		healPower = 200 - playerHealth
	}

	playerHealth += healPower

	return healPower
}

func AttackPlayer() int {
	attackPower := generateRandBetween(15, 30)

	playerHealth -= attackPower

	return attackPower
}

func GetHealthValues() (int, int) {
	return playerHealth, monsterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
