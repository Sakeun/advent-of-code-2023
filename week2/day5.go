package week2

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var input = "seeds: 79 14 55 13\n\n" +
	"seed-to-soil map:\n" + // 79 - 48 = 31 + 50 = soil 81.
	"50 98 2\n" +
	"52 50 48\n\n" +
	"soil-to-fertilizer map:\n" +
	"0 15 37\n" +
	"37 52 2\n" +
	"39 0 15\n\n" +
	"fertilizer-to-water map:\n" +
	"49 53 8\n" +
	"0 11 42\n" +
	"42 0 7\n" +
	"57 7 4\n\n" +
	"water-to-light map:\n" +
	"88 18 7\n" +
	"18 25 70\n\n" +
	"light-to-temperature map:\n" +
	"45 77 23\n" +
	"81 45 19\n" +
	"68 64 13\n\n" +
	"temperature-to-humidity map:\n" +
	"0 69 1\n1" +
	" 0 69\n\n" +
	"humidity-to-location map:\n" +
	"60 56 37\n" +
	"56 93 4"

func Day5Part1() int {
	inputList := strings.Split(string(getInput("week2/inputs/inputDay5.txt")), "\n\n")
	var values [][][]int
	for _, val := range inputList {
		var currenValues [][]int
		mapList := strings.Split(val, "\n")
		for _, list := range mapList {
			ints := parseInts(list)
			if len(ints) > 0 {
				currenValues = append(currenValues, ints)
			}
		}
		values = append(values, currenValues)
	}
	locs := values[0][0]
	values = values[1:]
	for _, mapValues := range values {
		for i, loc := range locs {
			for _, value := range mapValues {
				if loc > value[1] && loc < value[1]+value[2] {
					locs[i] = (loc - value[1]) + value[0]
				}
			}
		}
	}
	return slices.Min(locs)
}

func parseInts(str string) []int {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)

	var numArr []int
	for _, val := range nums {
		num, _ := strconv.Atoi(val)
		numArr = append(numArr, int(num))
	}
	return numArr
}

func getInput(input string) []byte {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
