package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var mapNames = []string{"soil", "fert", "water", "light", "temp", "humidity", "location"}

func main() {
	inputData := getInputData()
	answer := parseData(inputData)
	fmt.Println("Answer: {}", answer)
}

func parseData(inputData string) int {
	segments := strings.Split(inputData, "\n\n")
	seeds := parseSeeds(segments[0])
	rangeMaps := parseRangeMaps(segments[1:])
	return traverseMaps(seeds, rangeMaps)
}

func traverseMaps(seeds []int, maps [][]seedMapEntry) int {
	var smallestDistance = math.MaxInt
	for _, seed := range seeds {
		fmt.Printf("seed %d -> ", seed)
		var location = findLocation(seed, maps, 0)
		if location < smallestDistance {
			smallestDistance = location
		}
	}
	return smallestDistance
}

func findLocation(currentValue int, maps [][]seedMapEntry, depth int) int {
	var currentMap = maps[depth]
	for _, entry := range currentMap {
		if isInRange(currentValue, entry) {
			currentValue = currentValue - entry.sourceRangeStart + entry.destRangeStart
			break
		}
	}
	if depth+1 == len(maps) {
		fmt.Printf("%s %d\n", mapNames[depth], currentValue)
		return currentValue
	}
	fmt.Printf("%s %d -> ", mapNames[depth], currentValue)
	return findLocation(currentValue, maps, depth+1)
}

func isInRange(value int, mapEntry seedMapEntry) bool {
	return (value >= mapEntry.sourceRangeStart) && (value < mapEntry.sourceRangeStart+mapEntry.rangeLength)
}

func parseRangeMaps(segments []string) [][]seedMapEntry {
	maps := make([][]seedMapEntry, len(segments))
	for i, seg := range segments {
		maps[i] = parseMapSeg(seg)
	}
	return maps
}

func parseMapSeg(seg string) []seedMapEntry {
	var mapEntries []seedMapEntry
	entryLines := strings.Split(seg, "\n")

	for _, line := range entryLines[1:] {
		if line == "" {
			continue
		}
		partStrings := strings.Split(line, " ")
		parts := stringArrToIntArr(partStrings)
		entry := seedMapEntry{destRangeStart: parts[0], sourceRangeStart: parts[1], rangeLength: parts[2]}
		mapEntries = append(mapEntries, entry)
	}

	return mapEntries
}

func stringArrToIntArr(arr []string) []int {
	var ints []int

	for _, val := range arr {
		num, _ := strconv.ParseInt(val, 10, 32)
		ints = append(ints, int(num))
	}

	return ints
}

func parseSeeds(seedDef string) []int {
	data := strings.Replace(seedDef, "seeds: ", "", -1)
	seeds := strings.Split(data, " ")
	return stringArrToIntArr(seeds)
}

func getInputData() string {
	args := os.Args
	fPath := args[1]
	data, _ := os.ReadFile(fPath)
	content := string(data)
	return content
}

type seedMapEntry struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLength      int
}

const (
	seedsoil int = iota
	soilfert
	fertwater
	waterlight
	lighttemp
	temphumidity
	humiditylocation
)
