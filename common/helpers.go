package common

import (
	"log"
	"os"
)

func GetInput(input string) []byte {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
