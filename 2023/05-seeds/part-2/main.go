package main

import (
	"fmt"
	"os"
	"seeds/parsing"
)

func main() {
	inputData := getInputData()
	seeds, rangeMaps := parsing.ParsePuzzleDefinition(inputData)

	fmt.Print(seeds)
	fmt.Print(rangeMaps)
	fmt.Printf("Answer: %d", 0)
}

func getInputData() string {
	args := os.Args
	fPath := args[1]
	data, _ := os.ReadFile(fPath)
	content := string(data)
	return content
}
