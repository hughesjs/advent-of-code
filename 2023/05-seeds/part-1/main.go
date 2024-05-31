package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData := getInputData()
	answer := parseData(inputData)
	fmt.Printf("Answer: %d", answer)
}

func parseData(inputData string) int64 {
	segments := strings.Split(inputData, "\n\n")
	seeds := parseSeeds(segments[0])
	rangeMaps := parseRangeMaps(segments[1:])
	return traverseMaps(seeds, rangeMaps)
}

func traverseMaps(seeds []int64, maps [][]seedMapEntry) int64 {
	var smallestDistance int64 = math.MaxInt64
	for _, seed := range seeds {
		//fmt.Printf("seed %d -> ", seed)
		var location = findLocation(seed, maps, 0)
		if location < smallestDistance {
			smallestDistance = location
		}
	}
	return smallestDistance
}

func findLocation(currentValue int64, maps [][]seedMapEntry, depth int) int64 {
	var currentMap = maps[depth]
	for _, entry := range currentMap {
		if isInRange(currentValue, entry) {
			currentValue = currentValue - entry.sourceRangeStart + entry.destRangeStart
			break
		}
	}
	if depth+1 == len(maps) {
		//fmt.Printf("%s %d\n", mapNames[depth], currentValue)
		return currentValue
	}
	//fmt.Printf("%s %d -> ", mapNames[depth], currentValue)
	return findLocation(currentValue, maps, depth+1)
}

func isInRange(value int64, mapEntry seedMapEntry) bool {
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

func stringArrToIntArr(arr []string) []int64 {
	var ints []int64

	for _, val := range arr {
		num, _ := strconv.ParseInt(val, 10, 64)
		ints = append(ints, int64(num))
	}

	return ints
}

func parseSeeds(seedDef string) []int64 {
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
	destRangeStart   int64
	sourceRangeStart int64
	rangeLength      int64
}
