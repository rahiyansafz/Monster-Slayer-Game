package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())

var randGenerator = rand.New(randSource)

var currentMonstarHealth = MONSTAR_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonstar(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)

	currentMonstarHealth -= dmgValue

	return dmgValue

}

func HealPlayer() int {

	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() int {
	minAttackValue := MONSTAR_ATTACK_MIN_DMG
	maxAttackValue := MONSTAR_ATTACK_MAX_DMG

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)

	currentPlayerHealth -= dmgValue
	return dmgValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonstarHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
