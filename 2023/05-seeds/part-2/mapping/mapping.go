package mapping

import (
	"seeds/common"
	"slices"
)

func GetLocationRanges(startRanges []common.SeedRange, maps []common.SeedMap) []common.SeedRange {
	// TODO - Could recurse here to support arbitrary depth of maps
	soilRanges := applyMap(startRanges, maps[0])
	fertilizerRanges := applyMap(soilRanges, maps[1])
	waterRanges := applyMap(fertilizerRanges, maps[2])
	lightRanges := applyMap(waterRanges, maps[3])
	temperatureRanges := applyMap(lightRanges, maps[4])
	humidityRanges := applyMap(temperatureRanges, maps[5])
	locationRanges := applyMap(humidityRanges, maps[6])

	return locationRanges
}

func applyMap(startRanges []common.SeedRange, mapToApply common.SeedMap) []common.SeedRange {
	var newRanges []common.SeedRange

	for _, startRange := range startRanges {
		newRanges = slices.Concat(newRanges, getDestinationRanges(startRange, mapToApply))
	}
	return newRanges
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
	// Might be able to do this using floors and ceilings... I don't have a pencil and paper with me
	// TODO - Try this out later ^^

	// None of the range is in the search space
	if (startRange.End < inRange.Start) || (startRange.Start >= inRange.End) {
		return common.ZeroSeedRange(), common.ZeroSeedRange(), common.ZeroSeedRange()
	}

	// Whole of range in search space
	if inRange.ContainsAll(startRange) {
		return common.ZeroSeedRange(), startRange, common.ZeroSeedRange()
	}

	// Range completely subsumes search space
	if startRange.Start <= inRange.Start && startRange.End >= inRange.End {
		before := common.SeedRange{Start: startRange.Start, End: inRange.Start}
		after := common.SeedRange{Start: inRange.End, End: startRange.End}
		in := common.SeedRange{Start: inRange.Start, End: inRange.End}
		return before, in, after
	}

	// Range overlaps start of search space
	if (startRange.Start < inRange.Start) || inRange.Contains(startRange.End) {
		before := common.SeedRange{Start: startRange.Start, End: inRange.Start}
		in := common.SeedRange{Start: inRange.Start, End: startRange.End}
		after := common.ZeroSeedRange()
		return before, in, after
	}

	// Range overlaps end of search space
	if inRange.Contains(startRange.Start) || (startRange.End > inRange.End) {
		before := common.ZeroSeedRange()
		in := common.SeedRange{Start: startRange.Start, End: inRange.End}
		after := common.SeedRange{Start: inRange.End, End: startRange.End}
		return before, in, after
	}

	panic("This should be unreachable, if we get here, then I've fucked up")
}
