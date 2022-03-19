package main

import (
	"example/monstarSlayerGame/actions"
	"example/monstarSlayerGame/interaction"
)

var currentRound = 0

var gameRounds = []interaction.RoundData{}

func main() {

	startGame()

	winner := "" // "Player" || "Monstar" || ""

	for winner == "" {
		winner = excecuteRound()
	}

	endGame(winner)

}

func startGame() {
	interaction.PrintGreeting()
}

func excecuteRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealValue int
	var monstarAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonstar(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonstar(true)
	}

	monstarAttackDmg = actions.AttackPlayer()

	playerHealth, monstarHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonstarHealth:    monstarHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonstarAttackDmg: monstarAttackDmg,
	}

	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monstar"
	} else if monstarHealth <= 0 {
		return "Player"
	}

	return ""

}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
