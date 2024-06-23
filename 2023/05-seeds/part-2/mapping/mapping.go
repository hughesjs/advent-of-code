package mapping

import (
	"seeds/common"
	"slices"
)

func GetLocationRanges(startRanges []common.SeedRange, maps []common.SeedMap) []common.SeedRange {

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
	var destinationRanges []common.SeedRange

	for _, entry := range apply.Entries {
		if startRange.Start >= entry.SourceRange.End {
			continue
		}

		// TODO - These cases are broken

		// Whole of range in map
		if IsInRange(startRange.Start, entry.SourceRange) &&
			IsInRange(startRange.End, entry.SourceRange) {
			return []common.SeedRange{MapSubRangeToDestination(startRange, entry)}
		}

		// Start of range in the map
		if IsInRange(startRange.Start, entry.SourceRange) &&
			!IsInRange(startRange.End, entry.SourceRange) {
			subRangeInMap := common.SeedRange{Start: startRange.Start, End: entry.SourceRange.End}
			destinationRanges = append(destinationRanges, MapSubRangeToDestination(subRangeInMap, entry))

			subRangeOutOfMap := common.SeedRange{Start: entry.SourceRange.End, End: startRange.End}
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(subRangeOutOfMap, apply))
			return destinationRanges
		}

		// End of range in the map
		if !IsInRange(startRange.Start, entry.SourceRange) &&
			IsInRange(startRange.End, entry.SourceRange) {
			subRangeInMap := common.SeedRange{Start: entry.SourceRange.Start, End: startRange.End}
			destinationRanges = append(destinationRanges, MapSubRangeToDestination(subRangeInMap, entry))

			subRangeOutOfMap := common.SeedRange{Start: startRange.Start, End: entry.SourceRange.Start}
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(subRangeOutOfMap, apply))
			return destinationRanges
		}

		// Middle of range in the map
		if startRange.Start < entry.SourceRange.Start &&
			startRange.End > entry.SourceRange.End {
			subRangeInMap := common.SeedRange{Start: entry.SourceRange.Start, End: entry.SourceRange.End}
			destinationRanges = append(destinationRanges, MapSubRangeToDestination(subRangeInMap, entry))

			subRangeBeforeMap := common.SeedRange{Start: startRange.Start, End: entry.SourceRange.Start}
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(subRangeBeforeMap, apply))

			subRangeAfterMap := common.SeedRange{Start: entry.SourceRange.End, End: startRange.End}
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(subRangeAfterMap, apply))

			return destinationRanges
		}

	}

	// No matches
	return []common.SeedRange{startRange}
}

func MapSubRangeToDestination(subRange common.SeedRange, entry common.SeedMapEntry) common.SeedRange {
	if !(IsInRange(subRange.Start, entry.SourceRange) && subRange.End <= entry.SourceRange.End) {
		panic("Can't map subrange to destination unless it's in the source range")
	}

	offset := subRange.Start - entry.SourceRange.Start
	length := subRange.End - subRange.Start
	return common.SeedRange{Start: entry.DestRangeStart + offset, End: entry.DestRangeStart + offset + length}
}

// IsInRange Todo - Delete this and use the below
func IsInRange(value int64, sRange common.SeedRange) bool {
	return value >= sRange.Start && value < sRange.End
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
	if inRange.Contains(startRange.Start) && inRange.Contains(startRange.End) {
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
