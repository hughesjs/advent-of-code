package common

type SeedMapEntry struct {
	SourceRange    SeedRange
	DestRangeStart int64
}

// SeedRange [Start..End) - Includes start and excludes end
type SeedRange struct {
	Start int64
	End   int64
}

type SeedMap struct {
	Entries        []SeedMapEntry
	MapDescription string
}

func (s SeedRange) IsEmpty() bool { return s.Start == s.End }

func (s SeedRange) Contains(value int64) bool {
	return value >= s.Start && value < s.End
}

func ZeroSeedRange() SeedRange {
	return SeedRange{Start: 0, End: 0}
}
