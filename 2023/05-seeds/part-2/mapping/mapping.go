package mapping

import (
	"seeds/common"
	"slices"
)

func GetLocationRanges(startRanges []common.SeedRange, maps []common.SeedMap) []common.SeedRange {
	locationRanges := applyMap(startRanges, maps, 0)
	return locationRanges
}

func applyMap(startRanges []common.SeedRange, maps []common.SeedMap, depth int) []common.SeedRange {
	if depth >= len(maps) {
		return startRanges
	}

	var newRanges []common.SeedRange

	for _, startRange := range startRanges {
		newRanges = slices.Concat(newRanges, getDestinationRanges(startRange, maps[depth]))
	}

	return applyMap(newRanges, maps, depth+1)
}

func getDestinationRanges(startRange common.SeedRange, apply common.SeedMap) []common.SeedRange {

	// Arguably unnecessaru
	if startRange.IsEmpty() {
		panic("Start range cannot be empty")
	}

	var destinationRanges []common.SeedRange

	for _, entry := range apply.Entries {
		before, in, after := GetIntersectionAndRemainder(startRange, entry.SourceRange)

		if in.IsEmpty() {
			continue
		}

		destinationRanges = append(destinationRanges, MapSubRangeToDestination(in, entry))

		if !before.IsEmpty() {
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(before, apply))
		}

		if !after.IsEmpty() {
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(after, apply))
		}

		return destinationRanges
	}

	// No matches
	return []common.SeedRange{startRange}
}

func MapSubRangeToDestination(subRange common.SeedRange, entry common.SeedMapEntry) common.SeedRange {
	if !(entry.SourceRange.ContainsAll(subRange)) {
		panic("Can't map subrange to destination unless it's in the source range")
	}

	offset := subRange.Start - entry.SourceRange.Start
	length := subRange.End - subRange.Start
	return common.SeedRange{Start: entry.DestRangeStart + offset, End: entry.DestRangeStart + offset + length}
}

// GetIntersectionAndRemainder This finds the proportion of a range within another range and returns the coincident portion, the remainder before and the remainder after
func GetIntersectionAndRemainder(startRange common.SeedRange, inRange common.SeedRange) (beforeRange common.SeedRange, insideRange common.SeedRange, afterRange common.SeedRange) {

	// None of the range is in the search space
	if (startRange.End < inRange.Start) || (startRange.Start >= inRange.End) {
		return common.ZeroSeedRange(), common.ZeroSeedRange(), common.ZeroSeedRange()
	}

	before := common.SeedRange{Start: min(startRange.Start, inRange.Start), End: inRange.Start}
	in := common.SeedRange{Start: max(startRange.Start, inRange.Start), End: min(startRange.End, inRange.End)}
	after := common.SeedRange{Start: inRange.End, End: max(startRange.End, inRange.End)}

	return before, in, after
}
