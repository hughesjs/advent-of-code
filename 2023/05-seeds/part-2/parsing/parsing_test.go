package parsing

import (
	"github.com/stretchr/testify/assert"
	"seeds/common"
	"testing"
)

func TestParsePuzzleDefinition(t *testing.T) {
	const providedCase = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

	ranges, maps := ParsePuzzleDefinition(providedCase)

	assert.Len(t, ranges, 2)
	assert.Len(t, maps, 7)
}

func TestParseMapSeg(t *testing.T) {
	const testData = `seed-to-soil map:
50 98 2
52 50 48`

	expected := common.SeedMap{
		Entries: []common.SeedMapEntry{
			{SourceRange: common.SeedRange{Start: 98, End: 100}, DestRangeStart: 50},
			{SourceRange: common.SeedRange{Start: 50, End: 98}, DestRangeStart: 52},
		},
		MapDescription: "seed-to-soil",
	}

	results := parseMapSeg(testData)

	assert.Equal(t, expected, results)
}

func TestParseSeedRanges(t *testing.T) {
	const testData = "seeds: 79 14 55 13"

	expected := []common.SeedRange{
		{Start: 79, End: 93},
		{Start: 55, End: 68},
	}

	results := parseSeedRanges(testData)

	assert.Equal(t, expected, results)
}
