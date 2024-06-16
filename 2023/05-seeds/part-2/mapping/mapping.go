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

		// Whole of range in map
		if IsInRange(startRange.Start, entry.SourceRange) &&
			IsInRange(startRange.End, entry.SourceRange) {
			offset := startRange.Start - entry.SourceRange.Start
			length := startRange.End - startRange.Start
			destinationRange := common.SeedRange{Start: entry.DestRangeStart + offset, End: entry.DestRangeStart + offset + length}
			return []common.SeedRange{destinationRange}
		}

		// Start of range in the map
		if IsInRange(startRange.Start, entry.SourceRange) &&
			!IsInRange(startRange.End, entry.SourceRange) {
			subRangeInMap := common.SeedRange{Start: startRange.Start, End: entry.SourceRange.End}
			offset := subRangeInMap.Start - entry.SourceRange.Start
			length := startRange.End - startRange.Start
			destinationRange := common.SeedRange{Start: entry.DestRangeStart + offset, End: entry.DestRangeStart + offset + length}
			destinationRanges = append(destinationRanges, destinationRange)

			subRangeOutOfMap := common.SeedRange{Start: entry.SourceRange.End, End: startRange.End}
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(subRangeOutOfMap, apply))
			return destinationRanges
		}

		// End of range in the map
		if !IsInRange(startRange.Start, entry.SourceRange) &&
			IsInRange(startRange.End, entry.SourceRange) {
			subRangeInMap := common.SeedRange{Start: entry.SourceRange.Start, End: startRange.End}
			offset := subRangeInMap.Start - entry.SourceRange.Start
			length := startRange.End - startRange.Start
			destinationRange := common.SeedRange{Start: entry.DestRangeStart + offset, End: entry.DestRangeStart + offset + length}
			destinationRanges = append(destinationRanges, destinationRange)

			subRangeOutOfMap := common.SeedRange{Start: startRange.Start, End: entry.SourceRange.Start}
			destinationRanges = slices.Concat(destinationRanges, getDestinationRanges(subRangeOutOfMap, apply))
			return destinationRanges
		}

		// Middle of range in the map
		if startRange.Start < entry.SourceRange.Start &&
			startRange.End > entry.SourceRange.End {
			subRangeInMap := common.SeedRange{Start: entry.SourceRange.Start, End: entry.SourceRange.End}
			offset := subRangeInMap.Start - entry.SourceRange.Start
			length := startRange.End - startRange.Start
			destinationRange := common.SeedRange{Start: entry.DestRangeStart + offset, End: entry.DestRangeStart + offset + length}
			destinationRanges = append(destinationRanges, destinationRange)

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

func IsInRange(value int64, sRange common.SeedRange) bool {
	return value >= sRange.Start && value < sRange.End
}
