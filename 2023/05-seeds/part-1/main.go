package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData := getInputData()
	answer := parseData(inputData)
	fmt.Println("Answer: {}", answer)
}

func parseData(inputData string) int {
	segments := strings.Split(inputData, "\n\n")
	seeds := parseSeeds(segments[0])
	fmt.Println(seeds)
	rangeMaps := parseRangeMaps(segments[1:])
	fmt.Println(rangeMaps)
	expandedMaps := expandMaps(rangeMaps)
	fmt.Println(expandedMaps)
	return 0
}

func expandMaps(maps [][]seedMapEntry) []map[int]int {
	expandedMap := make([]map[int]int, len(maps))

	for i, mapEntries := range maps {
		expandedMap[i] = make(map[int]int)
		for _, mapEntry := range mapEntries {
			for k := 0; k < mapEntry.rangeLength; k++ {
				expandedMap[i][mapEntry.sourceRangeStart+k] = mapEntry.destRangeStart + k
			}
		}
	}
	return expandedMap
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
