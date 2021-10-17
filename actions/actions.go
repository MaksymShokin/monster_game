package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var MonsterHealth = MONSTER_STARTING_HEALTH
var PlayerHealth = PLAYER_STARTING_HEALTH

func AttackMonster(isSpecial bool) {
	var attackPower int
	if isSpecial {
		attackPower = generateRandBetween(10, 25)
	} else {
		attackPower = generateRandBetween(15, 35)
	}

	MonsterHealth -= attackPower
}

func HealPlayer() {
	healPower := generateRandBetween(25, 35)

	PlayerHealth += healPower

	if PlayerHealth > PLAYER_STARTING_HEALTH {
		PlayerHealth = PLAYER_STARTING_HEALTH
	}
}

func AttackPlayer() {
	attackPower := generateRandBetween(15, 30)

	PlayerHealth -= attackPower
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
