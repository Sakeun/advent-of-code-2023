package week1

import (
	"fmt"
	"github.com/Sakeun/advent-of-code-2023/common"
	"strconv"
	"strings"
)

var str = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
var charsMap = make(map[string]bool)

// 549908

func Day3Part1() int {
	input := common.GetInput("week1/inputs/inputDay3.txt")
	inputArr := strings.Split(string(input), "\n")
	sum = 0
	var sumGears int
	matrix := createMatrix(inputArr)

	for i, row := range matrix {
		if i == 0 || i == len(matrix)-1 {
			continue
		}
		for j, field := range row {
			if charsMap[field] {
				// Check left and right
				var err error
				_, err = strconv.Atoi(matrix[i][j-1])
				if err == nil {
					sum += parseInt(matrix[i][j-3], matrix[i][j-2], matrix[i][j-1], true)
				}

				_, err = strconv.Atoi(matrix[i][j+1])
				if err == nil {
					sum += parseInt(matrix[i][j+1], matrix[i][j+2], matrix[i][j+3], false)
				}

				// top
				checkAround(matrix[i-1], j)

				// Bottom, left bottom, right bottom
				checkAround(matrix[i+1], j)

				if field == "*" {
					sumGears += checkGear(matrix[i-1], matrix[i], matrix[i+1], j)
				}
			}
		}
	}
	fmt.Println(sumGears)
	return sum
}

func checkGear(top []string, mid []string, bot []string, j int) int {
	var int1, int2 int
	if _, err := strconv.Atoi(mid[j-1]); err == nil {
		num := parseInt(mid[j-3], mid[j-2], mid[j-1], true)
		int1 = num
	}
	if _, err := strconv.Atoi(mid[j+1]); err == nil {
		num := parseInt(mid[j+1], mid[j+2], mid[j+3], false)
		if int1 == 0 {
			int1 = num
		} else {
			int2 = num
		}
	}

	int1, int2 = checkGearRows(top, j, int1, int2)

	int1, int2 = checkGearRows(bot, j, int1, int2)

	return int1 * int2
}

func parseTB(a string, b string, c string, d string, e string) int {
	returnVal := ""
	if a != "." && !charsMap[a] {
		returnVal += a
	}
	if b != "." && !charsMap[b] {
		returnVal += b
	} else {
		returnVal = ""
	}
	returnVal += c
	if d != "." && !charsMap[d] {
		returnVal += d
		if e != "." && !charsMap[e] {
			returnVal += e
		}
	}
	num, _ := strconv.Atoi(returnVal)
	return num
}

func parseInt(a string, b string, c string, isLeft bool) int {
	returnVal := ""
	if a != "." && !charsMap[a] {
		returnVal += a
	}

	if b != "." && !charsMap[b] {
		returnVal += b
	} else {
		if !isLeft {
			num, _ := strconv.Atoi(returnVal)
			return num
		}
		returnVal = ""
	}

	if c != "." && !charsMap[c] {
		returnVal += c
	}
	num, _ := strconv.Atoi(returnVal)
	return num
}

func checkGearRows(row []string, i int, a int, b int) (int, int) {
	if _, err := strconv.Atoi(row[i]); err == nil {
		if a != 0 && b != 0 {
			return 0, 0
		}
		num := parseTB(row[i-2], row[i-1], row[i], row[i+1], row[i+2])
		if a == 0 {
			a = num
		} else {
			b = num
		}
	} else {
		if _, err = strconv.Atoi(row[i+1]); err == nil {
			if a != 0 && b != 0 {
				return 0, 0
			}
			num := parseInt(row[i+1], row[i+2], row[i+3], false)
			if a == 0 {
				a = num
			} else {
				b = num
			}
		}
		if _, err = strconv.Atoi(row[i-1]); err == nil {
			if a != 0 && b != 0 {
				return 0, 0
			}
			num := parseInt(row[i-3], row[i-2], row[i-1], true)
			if a == 0 {
				a = num
			} else {
				b = num
			}
		}
	}
	return a, b
}

func checkAround(row []string, i int) {
	if _, err := strconv.Atoi(row[i]); err == nil {
		sum += parseTB(row[i-2], row[i-1], row[i], row[i+1], row[i+2])
	} else {
		_, err = strconv.Atoi(row[i-1])
		if err == nil {
			sum += parseInt(row[i-3], row[i-2], row[i-1], true)
		}
		_, err = strconv.Atoi(row[i+1])
		if err == nil {
			sum += parseInt(row[i+1], row[i+2], row[i+3], false)
		}
	}
}

func createMatrix(input []string) [][]string {
	var matrix [][]string
	for _, rows := range input {
		var row []string
		for _, field := range rows {
			strfield := string(field)
			row = append(row, strfield)
			if _, err := strconv.Atoi(strfield); err != nil && strfield != "." {
				charsMap[strfield] = true
			}
		}
		matrix = append(matrix, row)
	}
	return matrix
}
