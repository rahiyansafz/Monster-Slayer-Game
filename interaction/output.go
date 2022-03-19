package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonstarAttackDmg int
	PlayerHealth     int
	MonstarHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTAR SLAYER", "", true)
	asciiFigure.Print()
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!!!")
}

func ShowAvailableActions(specialIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monstar")
	fmt.Println("(2) Heal")

	if specialIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monstar for %v damage\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monstar for %v damage\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v damage\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monstar attacked player for %v damage\n", roundData.MonstarAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monstar Health: %v\n", roundData.MonstarHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	asciiFigure := figure.NewColorFigure("GAME OVER!", "", "red", true)
	asciiFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {

	exPath, err := os.Executable()
	if err != nil {
		fmt.Println("Writing log file failed. Exiting.")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")
	// file, err := os.Create("gamelog.txt") // Required for go run .

	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monstar Attack Damage": fmt.Sprint(value.MonstarAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monstar Health":        fmt.Sprint(value.MonstarHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting.")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log!")

}
