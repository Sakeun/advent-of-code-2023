package week1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var firstNum string
var secondNum string
var sum int

func Day1Part1() int {
	sum, firstNum, secondNum = 0, "", ""
	content := getInput()
	for _, val := range string(content) {
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
	content := getInput()
	arr := strings.Fields(string(content))
	for _, val := range arr {
		for i, r := range val {
			currNum := ""
			switch {
			case i+2 < len(val) && string(r) == "o" && val[i:i+3] == "one":
				currNum = "1"
			case string(r) == "t":
				if i+2 < len(val) && val[i:i+3] == "two" {
					currNum = "2"
				} else if i+4 < len(val) && val[i:i+5] == "three" {
					currNum = "3"
				}
			case string(r) == "f":
				if i+3 < len(val) && val[i:i+4] == "four" {
					currNum = "4"
				} else if i+3 < len(val) && val[i:i+4] == "five" {
					currNum = "5"
				}
			case string(r) == "s":
				if i+4 < len(val) && val[i:i+5] == "seven" {
					currNum = "7"
				} else if i+2 < len(val) && val[i:i+3] == "six" {
					currNum = "6"
				}
			case i+4 < len(val) && string(r) == "e" && val[i:i+5] == "eight":
				currNum = "8"
			case i+3 < len(val) && string(r) == "n" && val[i:i+4] == "nine":
				currNum = "9"
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

func getInput() []byte {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return content
}
