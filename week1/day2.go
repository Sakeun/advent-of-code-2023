package week1

import (
	"github.com/Sakeun/advent-of-code-2023/common"
	"strconv"
	"strings"
)

func Day2Part1() int {
	firstNum, secondNum, sum = "", "", 0
	file := common.GetInput("week1/inputs/inputDay2.txt")
	for _, val := range file[:len(file)-1] {
		var game int
		gameId, gameSets, _ := strings.Cut(val, ": ")
		game, _ = strconv.Atoi(strings.Fields(gameId)[1])

		for _, allRounds := range strings.Split(gameSets, "; ") {
			rounds := strings.Split(allRounds, ", ")
			var red, blue, green int

			for _, picks := range rounds {
				color := strings.Split(picks, " ")
				currGame, _ := strconv.Atoi(color[0])
				switch color[1] {
				case "red":
					red += currGame
				case "green":
					green += currGame
				case "blue":
					blue += currGame
				}
			}
			if red > 12 || green > 13 || blue > 14 {
				game = 0
				firstNum, secondNum = "", ""
				break
			}
		}
		sum += game
		firstNum, secondNum = "", ""
	}

	return sum
}

func Day2Part2() int {
	firstNum, secondNum, sum = "", "", 0
	file := common.GetInput("week1/inputs/inputDay2.txt")
	for _, val := range file[:len(file)-1] {
		var red, blue, green int
		_, gameSets, _ := strings.Cut(val, ": ")

		for _, allRounds := range strings.Split(gameSets, "; ") {
			rounds := strings.Split(allRounds, ", ")

			for _, picks := range rounds {
				color := strings.Split(picks, " ")
				currGame, _ := strconv.Atoi(color[0])
				switch {
				case color[1] == "red" && currGame > red:
					red = currGame
				case color[1] == "green" && currGame > green:
					green = currGame
				case color[1] == "blue" && currGame > blue:
					blue = currGame
				}
			}
		}
		sum += red * blue * green
		firstNum, secondNum = "", ""
	}

	return sum
}
