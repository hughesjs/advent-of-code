package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestExampleCase(t *testing.T) {
	expected := int(35)
	result := parseData(providedCase)
	assert.Equal(t, expected, result)
}

func TestMapExpansion(t *testing.T) {
	const providedCase = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48`

	seedSoil := expandMaps(parseRangeMaps(strings.Split(providedCase, "\n\n")[1:]))[0]

	assert.Len(t, seedSoil, 50)
	assert.Equal(t, seedSoil[50], 52)
	assert.Equal(t, seedSoil[51], 53)
	assert.Equal(t, seedSoil[96], 98)
	assert.Equal(t, seedSoil[97], 99)
	assert.Equal(t, seedSoil[98], 50)
	assert.Equal(t, seedSoil[99], 51)
}
