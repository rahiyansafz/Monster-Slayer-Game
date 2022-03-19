package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	PlayerChoiceAttack        = 1
	PlayerChoiceHeal          = 2
	playerChoiceSpecialAttack = 3
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == fmt.Sprint(PlayerChoiceAttack) {
			return "ATTACK"
		} else if playerChoice == fmt.Sprint(PlayerChoiceHeal) {
			return "HEAL"
		} else if playerChoice == fmt.Sprint(playerChoiceSpecialAttack) && specialIsAvailable {
			return "SPECIAL ATTACK"
		}
		fmt.Println("Fetching the player input failed. Please try again.")

		// if err != nil {
		// 	fmt.Println("Fetching the player input failed. Please try again.")
		// }
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	playerInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	playerInput = strings.Replace(playerInput, "\r\n", "", -1)

	return playerInput, nil
}
