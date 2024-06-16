package main

import (
	"fmt"
	"math"
	"os"
	"seeds/mapping"
	"seeds/parsing"
)

func main() {
	inputData := getInputData()
	answer := runPuzzle(inputData)
	fmt.Printf("Answer: %d", answer)
}

func runPuzzle(inputData string) int64 {
	seeds, rangeMaps := parsing.ParsePuzzleDefinition(inputData)
	locationRanges := mapping.GetLocationRanges(seeds, rangeMaps)
	lowestLocation := int64(math.MaxInt64)

	for _, r := range locationRanges {
		if r.Start < lowestLocation {
			lowestLocation = r.Start
		}
	}

	return lowestLocation
}

func getInputData() string {
	args := os.Args
	fPath := args[1]
	data, _ := os.ReadFile(fPath)
	content := string(data)
	return content
}
