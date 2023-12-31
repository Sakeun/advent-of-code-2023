package week1

import (
	"github.com/Sakeun/advent-of-code-2023/common"
	"log"
	"os"
	"strconv"
)

var firstNum string
var secondNum string
var sum int

func Day1Part1() int {
	sum, firstNum, secondNum = 0, "", ""
	content, _ := os.ReadFile("week1/inputs/inputDay1.txt")
	for _, val := range content {
		if val > 48 && val <= 57 {
			if firstNum == "" {
				firstNum = string(val)
				secondNum = string(val)
			} else {
				secondNum = string(val)
			}
		}
		if val == 10 {
			total := firstNum + secondNum
			num, err := strconv.Atoi(total)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
			firstNum, secondNum = "", ""
		}
	}

	return sum
}

func Day1Part2() int {
	sum, firstNum, secondNum = 0, "", ""
	content := common.GetInput("week1/inputs/inputDay1.txt")
	content = content[:len(content)-1]
	for _, val := range content {
		for i, r := range val {
			currNum := ""
			switch {
			case string(r) == "o":
				currNum = findNum("one", val[i:])
			case string(r) == "t":
				currNum = findNum("two", val[i:])
				if currNum == "" {
					currNum = findNum("three", val[i:])
				}
			case string(r) == "f":
				currNum = findNum("four", val[i:])
				if currNum == "" {
					currNum = findNum("five", val[i:])
				}
			case string(r) == "s":
				currNum = findNum("six", val[i:])
				if currNum == "" {
					currNum = findNum("seven", val[i:])
				}
			case string(r) == "e":
				currNum = findNum("eight", val[i:])
			case string(r) == "n":
				currNum = findNum("nine", val[i:])
			case r > 48 && r <= 57:
				currNum = string(r)
			}
			if firstNum == "" {
				firstNum = currNum
				secondNum = currNum
			} else if currNum != "" {
				secondNum = currNum
			}
		}
		total := firstNum + secondNum
		num, err := strconv.Atoi(total)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
		firstNum, secondNum = "", ""
	}

	return sum
}

func findNum(num string, val string) string {
	vals := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if len(num)-1 < len(val) && val[:len(num)] == num {
		return vals[num]
	}
	return ""
}
