package week1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() int {
	firstNum, secondNum, sum = "", "", 0
	file := getInput("inputDay2.txt")
	arr := strings.Split(string(file), "\n")
	for j, val := range arr {
		var red, green, blue, game int
		for i, c := range val {
			if string(c) == ";" {
				red, green, blue = 0, 0, 0
				continue
			}
			if c > 48 && c <= 57 && firstNum == "" {
				firstNum = string(c)
				continue
			}
			if c >= 48 && c <= 57 {
				secondNum = string(c)
				continue
			}
			if string(c) == ":" {
				game = getNum(firstNum, secondNum)
				if j == 99 {
					game += 90
				}
				firstNum, secondNum = "", ""
				continue
			}
			if string(c) == "b" && val[i:i+4] == "blue" {
				blue += getNum(firstNum, secondNum)
				firstNum, secondNum = "", ""
			}
			if string(c) == "g" && val[i:i+5] == "green" {
				green += getNum(firstNum, secondNum)
				firstNum, secondNum = "", ""
			}
			if string(c) == "r" && val[i:i+3] == "red" {
				red += getNum(firstNum, secondNum)
				firstNum, secondNum = "", ""
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
	file := getInput("inputDay2.txt")
	arr := strings.Split(string(file), "\n")
	for _, val := range arr {
		var red, green, blue, power int
		for i, c := range val {
			var num int
			if c > 48 && c <= 57 && firstNum == "" {
				firstNum = string(c)
				continue
			}
			if c >= 48 && c <= 57 {
				secondNum = string(c)
				continue
			}
			if firstNum != "" {
				num = getNum(firstNum, secondNum)
			}
			if string(c) == ":" {
				firstNum, secondNum = "", ""
				continue
			}
			if string(c) == "b" && val[i:i+4] == "blue" {
				if num > blue {
					blue = num
				}
				firstNum, secondNum = "", ""
			}
			if string(c) == "g" && val[i:i+5] == "green" {
				if num > green {
					green = num
				}
				firstNum, secondNum = "", ""
			}
			if string(c) == "r" && val[i:i+3] == "red" {
				if num > red {
					red = num
				}
				firstNum, secondNum = "", ""
			}
		}
		power = red * blue * green
		sum += power
		firstNum, secondNum = "", ""
	}

	return sum
}

func getNum(first string, second string) int {
	total := first + second
	num, err := strconv.Atoi(total)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func getInput(input string) []byte {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
