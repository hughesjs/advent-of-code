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

func parseData(inputData string) int32 {
	segments := strings.Split(inputData, "\n\n")
	seeds := parseSeeds(segments[0])
	fmt.Println(seeds)
	maps := parseMaps(segments[1:])
	fmt.Println(maps)
	return 0
}

func parseMaps(segments []string) map[int][]seedMapEntry {
	maps := make(map[int][]seedMapEntry)
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

func stringArrToIntArr(arr []string) []int32 {
	var ints []int32

	for _, val := range arr {
		num, _ := strconv.ParseInt(val, 10, 32)
		ints = append(ints, int32(num))
	}

	return ints
}

func parseSeeds(seedDef string) []int32 {
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
	destRangeStart   int32
	sourceRangeStart int32
	rangeLength      int32
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
