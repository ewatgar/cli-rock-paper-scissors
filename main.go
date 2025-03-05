package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/ewatgar/cli-rock-paper-scissors/option"
	"github.com/ewatgar/cli-rock-paper-scissors/status"
	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose 'rock', 'paper' or 'scissors'! ")
		color.White("('quit' to stop playing):")
		text, err := reader.ReadString('\n')
		text = strings.ToLower(strings.Trim(text, " \r\n"))

		if text == "quit" {
			fmt.Println("Bye!")
			return
		}

		if err != nil {
			color.Red(fmt.Sprint("CRASH!! - ERROR:", err))
			return
		}

		playerOption, err := option.FromString(text)

		if err != nil {
			fmt.Println("Invalid option. Try again")
			continue
		}

		rivalOption := option.GameOption(rand.Intn(3))
		battleStatus := status.CheckWinStatus(playerOption, rivalOption)

		switch battleStatus {
		case status.Win:
			color.Green(fmt.Sprint("You won! ", playerOption, " beats ", rivalOption))
			return
		case status.Lose:
			color.Red(fmt.Sprint("You lost... ", rivalOption, " beats ", playerOption))
			return
		case status.Stalemate:
			color.Yellow("Stalemate! Try again")
		}
	}
}
