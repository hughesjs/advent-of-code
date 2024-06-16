package parsing

import (
	"seeds/common"
	"strconv"
	"strings"
)

func ParsePuzzleDefinition(inputData string) ([]common.SeedRange, []common.SeedMap) {
	segments := strings.Split(inputData, "\n\n")
	seedRanges := parseSeedRanges(segments[0])
	rangeMaps := parseRangeMaps(segments[1:])

	return seedRanges, rangeMaps
}

func parseRangeMaps(segments []string) []common.SeedMap {
	maps := make([]common.SeedMap, len(segments))
	for i, seg := range segments {
		maps[i] = parseMapSeg(seg)
	}
	return maps
}

func parseMapSeg(seg string) common.SeedMap {
	var mapEntries []common.SeedMapEntry
	entryLines := strings.Split(seg, "\n")

	for _, line := range entryLines[1:] {
		if line == "" {
			continue
		}
		partStrings := strings.Split(line, " ")
		parts := stringArrToIntArr(partStrings)
		entry := common.SeedMapEntry{SourceRange: common.SeedRange{Start: parts[1], End: parts[1] + parts[2]}, DestRangeStart: parts[0]}
		mapEntries = append(mapEntries, entry)
	}
	name := strings.Replace(entryLines[0], " map:", "", -1)
	return common.SeedMap{Entries: mapEntries, MapDescription: name}
}

func stringArrToIntArr(arr []string) []int64 {
	var ints []int64

	for _, val := range arr {
		num, _ := strconv.ParseInt(val, 10, 64)
		ints = append(ints, int64(num))
	}

	return ints
}

func parseSeedRanges(seedDef string) []common.SeedRange {
	data := strings.Replace(seedDef, "seeds: ", "", -1)
	seeds := strings.Split(data, " ")
	rawInts := stringArrToIntArr(seeds)
	ranges := make([]common.SeedRange, len(rawInts)/2)

	for i := 0; i < len(rawInts); i += 2 {
		ranges[i/2] = common.SeedRange{Start: rawInts[i], End: rawInts[i] + rawInts[i+1]}
	}

	return ranges
}
