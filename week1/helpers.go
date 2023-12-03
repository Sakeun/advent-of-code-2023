package week1

import (
	"log"
	"os"
)

func getInput(input string) []byte {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
