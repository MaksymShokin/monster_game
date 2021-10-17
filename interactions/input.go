package interactions

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(hasSpecialAttack bool) string {
	for {
		choice, _ := getPlayerInput()

		if choice == "1" {
			return "ATTACK"
		} else if choice == "2" {
			return "HEAL"
		} else if choice == "3" && hasSpecialAttack {
			return "SPECIAL"
		}

		fmt.Println("Invalid input, please try again")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")
	data, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	data = parseInput(data)

	return data, nil
}

func parseInput(data string) string {
	if runtime.GOOS == "windows" {
		data = strings.Replace(data, "\r\n", "", -1)
	} else {
		data = strings.Replace(data, "\n", "", -1)
	}

	return data
}
