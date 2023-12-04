package week2

import (
	"github.com/Sakeun/advent-of-code-2023/common"
	"slices"
	"strconv"
	"strings"
)

var example = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"

func Day4Part1() int {
	var sum int
	input := common.GetInput("week2/inputs/inputDay4.txt")
	for _, card := range input {
		_, nums, _ := strings.Cut(card, ": ")
		left, right, _ := strings.Cut(nums, " | ")
		leftNumbers := strings.Fields(left)
		rightNumbers := strings.Fields(right)

		var points int

		for _, num := range leftNumbers {
			if slices.Contains(rightNumbers, num) {
				if points == 0 {
					points += 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}

	return sum
}

func Day4Part2() int {
	cardCount := make(map[int]int)
	var sum int
	input := common.GetInput("week2/inputs/inputDay4.txt")
	input = input[:len(input)-1]
	for _, card := range input {
		cardNum, nums, _ := strings.Cut(card, ": ")
		parsedNum, _ := strconv.Atoi(strings.Fields(cardNum)[1])
		cardCount[parsedNum] += 1
		left, right, _ := strings.Cut(nums, " | ")
		leftNumbers := strings.Fields(left)
		rightNumbers := strings.Fields(right)

		k := parsedNum
		for _, num := range leftNumbers {
			if slices.Contains(rightNumbers, num) {
				k += 1
				cardCount[k] += cardCount[parsedNum]
			}
		}
		sum += cardCount[parsedNum]
	}

	return sum
}
