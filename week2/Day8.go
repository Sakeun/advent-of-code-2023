package week2

import (
	"fmt"
	"slices"
	"strings"
)

type day8Node struct {
	left  string
	right string
}

var day8Input = "RL\n\n" +
	"AAA = (BBB, CCC)\n" +
	"BBB = (DDD, EEE)\n" +
	"CCC = (ZZZ, GGG)\n" +
	"DDD = (DDD, DDD)\n" +
	"EEE = (EEE, EEE)\n" +
	"GGG = (GGG, GGG)\n" +
	"ZZZ = (ZZZ, ZZZ)\n"

var day8Input2 = "LR\n\n" +
	"11A = (11B, XXX)\n" +
	"11B = (XXX, 11Z)\n" +
	"11Z = (11B, XXX)\n" +
	"22A = (22B, XXX)\n" +
	"22B = (22C, 22C)\n" +
	"22C = (22Z, 22Z)\n" +
	"22Z = (22B, 22B)\n" +
	"XXX = (XXX, XXX)"

func Day8Part1() int {
	paths := strings.Split(string(getInput("week2/inputs/inputDay8.txt")), "\n\n")
	instructions := paths[0]
	allPaths := strings.Split(paths[1], "\n")
	allPaths = allPaths[:len(allPaths)-1]
	pathMap := make(map[string]day8Node)
	for _, path := range allPaths {
		name := strings.Split(path, " = (")
		leftRight := strings.Split(name[1], ", ")
		leftRight[1] = leftRight[1][:len(leftRight[1])-1]
		pathMap[name[0]] = day8Node{leftRight[0], leftRight[1]}
	}
	node := "AAA"
	var steps int
	for {
		for _, inst := range instructions {
			steps++
			if string(inst) == "R" {
				node = pathMap[node].right
			} else {
				node = pathMap[node].left
			}
			if node == "ZZZ" {
				break
			}
		}
		if node == "ZZZ" {
			break
		}
	}
	return steps
}

func Day8Part2() int {
	paths := strings.Split(string(getInput("week2/inputs/inputDay8.txt")), "\n\n")
	//paths := strings.Split(day8Input2, "\n\n")
	instructions := paths[0]
	allPaths := strings.Split(paths[1], "\n")
	allPaths = allPaths[:len(allPaths)-1]
	pathMap := make(map[string]day8Node)

	var allStartNodes []string
	for _, path := range allPaths {
		name := strings.Split(path, " = (")
		leftRight := strings.Split(name[1], ", ")
		leftRight[1] = leftRight[1][:len(leftRight[1])-1]
		pathMap[name[0]] = day8Node{leftRight[0], leftRight[1]}
		if string(name[0][2]) == "A" {
			allStartNodes = append(allStartNodes, name[0])
		}
	}
	var allDurations []int
	for _, node := range allStartNodes {
		allDurations = append(allDurations, getPath(node, pathMap, instructions))
	}
	steps := 1
	fmt.Println(allDurations)
	return leastCommonMultiple(allDurations)
	//for {
	//	steps++
	//	allZero := false
	//	for _, duration := range allDurations {
	//		if steps%duration != 0 {
	//			allZero = false
	//			break
	//		} else {
	//			fmt.Println("Zero")
	//			allZero = true
	//		}
	//	}
	//	if allZero {
	//		break
	//	}
	//}
	return steps
}

func leastCommonMultiple(nums []int) int {
	highest := slices.Max(nums)
	allFound := false
	for i := 1; true; i++ {
		currMultiple := highest * i
		for _, val := range nums {
			if currMultiple%val != 0 {
				allFound = false
				break
			} else {
				allFound = true
			}
		}
		if allFound {
			return currMultiple
		}
	}
	return 0
}

func getPath(node string, nodeMap map[string]day8Node, instructions string) int {
	var steps int
	for {
		for _, inst := range instructions {
			steps++
			if string(inst) == "R" {
				node = nodeMap[node].right
			} else {
				node = nodeMap[node].left
			}
			if string(node[2]) == "Z" {
				break
			}
		}
		if string(node[2]) == "Z" {
			break
		}
	}

	return steps
}
