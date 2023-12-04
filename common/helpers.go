package common

import (
	"log"
	"os"
	"strings"
)

func GetInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	inputArr := strings.Split(string(content), "\n")
	return inputArr
}
