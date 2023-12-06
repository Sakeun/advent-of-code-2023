package week2

import (
	"strings"
)

var inputDay6 = "Time:      7  15   30\n" +
	"Distance:  9  40  200"

func Day6Part1() int {
	values := strings.Split(string(getInput("week2/inputs/inputDay6.txt")), "\n")
	times := parseInts(values[0])
	distances := parseInts(values[1])

	totalTime := 1
	for j, time := range times {
		amount := 0
		for i := 0; i < time; i++ {
			if i*(time-i) > distances[j] {
				amount++
			}
		}
		totalTime *= amount
	}
	return totalTime
}

func Day6Part2() int {
	values := strings.Split(string(getInput("week2/inputs/inputDay6.txt")), "\n")

	_, tempTimes, _ := strings.Cut(values[0], ":")
	_, tempDistances, _ := strings.Cut(values[1], ":")

	times := strings.Fields(tempTimes)
	distances := strings.Fields(tempDistances)

	var timeString string
	var distanceString string
	for i, val := range times {
		timeString = timeString + val
		distanceString = distanceString + distances[i]
	}
	distance := parseInts(distanceString)[0]
	time := parseInts(timeString)[0]
	totalTime := 0
	for i := 0; i < time; i++ {
		if totalTime > 1 && i*(time-i) < distance {
			break
		}
		if i*(time-i) > distance {
			totalTime++
		}
	}
	return totalTime
}
