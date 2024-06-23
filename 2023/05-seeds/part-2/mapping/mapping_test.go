package mapping

import (
	"github.com/stretchr/testify/assert"
	"seeds/common"
	"testing"
)

func TestMapSubRangeToDestination(t *testing.T) {
	testEntry := common.SeedMapEntry{SourceRange: common.SeedRange{Start: 1, End: 10}, DestRangeStart: 15}
	testSubRange := common.SeedRange{Start: 5, End: 10}
	expected := common.SeedRange{Start: 19, End: 24}

	result := MapSubRangeToDestination(testSubRange, testEntry)

	assert.Equal(t, expected, result)
}

func TestMapSubRangeToDestinationPanicIfNotInSourceRange(t *testing.T) {
	testEntry := common.SeedMapEntry{SourceRange: common.SeedRange{Start: 1, End: 10}, DestRangeStart: 15}
	testSubRange := common.SeedRange{Start: 5, End: 11}
	testSubRange2 := common.SeedRange{Start: 0, End: 1}
	testSubRange3 := common.SeedRange{Start: 0, End: 11}

	assert.Panics(t, func() { MapSubRangeToDestination(testSubRange, testEntry) })
	assert.Panics(t, func() { MapSubRangeToDestination(testSubRange2, testEntry) })
	assert.Panics(t, func() { MapSubRangeToDestination(testSubRange3, testEntry) })
}

func TestMappingWholeRangeInSingleEntry(t *testing.T) {
	testEntryOne := common.SeedMapEntry{SourceRange: common.SeedRange{Start: 50, End: 98}, DestRangeStart: 52}
	testEntryTwo := common.SeedMapEntry{SourceRange: common.SeedRange{Start: 98, End: 100}, DestRangeStart: 50}
	testMap := common.SeedMap{Entries: []common.SeedMapEntry{testEntryTwo, testEntryOne}, MapDescription: "test"}

	testRangeOne := common.SeedRange{Start: 55, End: 68}
	testRangeTwo := common.SeedRange{Start: 79, End: 93}

	expectedOne := []common.SeedRange{{Start: 57, End: 70}}
	expectedTwo := []common.SeedRange{{Start: 81, End: 95}}

	resultOne := getDestinationRanges(testRangeOne, testMap)
	resultTwo := getDestinationRanges(testRangeTwo, testMap)

	assert.Equal(t, expectedOne, resultOne)
	assert.Equal(t, expectedTwo, resultTwo)
}

func TestGetSubrangeWhenNoneWithinSearchRange(t *testing.T) {
	inRange := common.SeedRange{Start: 0, End: 10}
	searchRange := common.SeedRange{Start: 10, End: 15}

	before, in, after := GetIntersectionAndRemainder(searchRange, inRange)

	assert.Equal(t, before, common.ZeroSeedRange())
	assert.Equal(t, in, common.ZeroSeedRange())
	assert.Equal(t, after, common.ZeroSeedRange())
}

func TestGetSubrangeWhenAllWithinSearchRange(t *testing.T) {
	// [0..10)
	inRange := common.SeedRange{Start: 0, End: 10}
	searchRange := common.SeedRange{Start: 5, End: 7}

	expected := common.SeedRange{Start: 5, End: 7}

	before, in, after := GetIntersectionAndRemainder(searchRange, inRange)

	assert.True(t, before.IsEmpty())
	assert.True(t, after.IsEmpty())
	assert.Equal(t, expected, in)
}

func TestGetSubrangeWhenSubsumesSearchRange(t *testing.T) {
	inRange := common.SeedRange{Start: 4, End: 7}
	searchRange := common.SeedRange{Start: 0, End: 10}

	expectedBefore := common.SeedRange{Start: 0, End: 4}
	expectedIn := common.SeedRange{Start: 4, End: 7}
	expectedAfter := common.SeedRange{Start: 7, End: 10}

	before, in, after := GetIntersectionAndRemainder(searchRange, inRange)

	assert.Equal(t, expectedBefore, before)
	assert.Equal(t, expectedIn, in)
	assert.Equal(t, expectedAfter, after)
}

func TestGetSubrangeWhenOverlappingWithStartOfSearchRange(t *testing.T) {
	inRange := common.SeedRange{Start: 5, End: 10}
	searchRange := common.SeedRange{Start: 0, End: 7}

	expectedBefore := common.SeedRange{Start: 0, End: 5}
	expectedIn := common.SeedRange{Start: 5, End: 7}

	before, in, after := GetIntersectionAndRemainder(searchRange, inRange)

	assert.Equal(t, expectedBefore, before)
	assert.Equal(t, expectedIn, in)
	assert.True(t, after.IsEmpty())
}

func TestGetSubrangeWhenOverlappingWithEndOfSearchRange(t *testing.T) {
	inRange := common.SeedRange{Start: 0, End: 7}
	searchRange := common.SeedRange{Start: 5, End: 10}

	expectedBefore := common.ZeroSeedRange()
	expectedIn := common.SeedRange{Start: 5, End: 7}
	expectedAfter := common.SeedRange{Start: 7, End: 10}

	before, in, after := GetIntersectionAndRemainder(searchRange, inRange)

	assert.Equal(t, expectedBefore, before)
	assert.Equal(t, expectedIn, in)
	assert.Equal(t, expectedAfter, after)
}
