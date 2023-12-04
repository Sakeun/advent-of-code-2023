package week1

import (
	"fmt"
	"github.com/Sakeun/advent-of-code-2023/common"
	"regexp"
	"strconv"
)

var str = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
var charsMap = make(map[string]bool)

func Day3Part1() int {
	input := common.GetInput("week1/inputs/inputDay3.txt")
	sum = 0
	var num int
	var surroundingNums []int
	var sumGears int
	matrix := createMatrix(input)

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
					num = parseInt(matrix[i], j-3, true)
					sum += num
					surroundingNums = append(surroundingNums, num)
				}

				_, err = strconv.Atoi(matrix[i][j+1])
				if err == nil {
					num = parseInt(matrix[i], j+1, false)
					sum += num
					surroundingNums = append(surroundingNums, num)
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
		num := parseInt(mid, j-3, true)
		int1 = num
	}
	if _, err := strconv.Atoi(mid[j+1]); err == nil {
		num := parseInt(mid, j+1, false)
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

func parseInt(row []string, i int, isLeft bool) int {
	str := row[i] + row[i+1] + row[i+2]
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)

	if !isLeft || len(nums) == 1 {
		num, _ := strconv.Atoi(nums[0])
		return num
	}
	num, _ := strconv.Atoi(nums[1])
	return num
}

func parseTB(row []string, i int) int {
	str := row[i] + row[i+1] + row[i+2] + row[i+3] + row[i+4]
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)

	if len(nums) == 1 || (row[i+3] == "." || charsMap[row[i+3]]) {
		num, _ := strconv.Atoi(nums[0])
		return num
	}
	num, _ := strconv.Atoi(nums[1])
	return num
}

func checkGearRows(row []string, i int, a int, b int) (int, int) {
	if _, err := strconv.Atoi(row[i]); err == nil {
		if a != 0 && b != 0 {
			return 0, 0
		}
		num := parseTB(row, i-2)
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
			num := parseInt(row, i+1, false)
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
			num := parseInt(row, i-3, true)
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
		sum += parseTB(row, i-2)
	} else {
		_, err = strconv.Atoi(row[i-1])
		if err == nil {
			sum += parseInt(row, i-3, true)
		}
		_, err = strconv.Atoi(row[i+1])
		if err == nil {
			sum += parseInt(row, i+1, false)
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
