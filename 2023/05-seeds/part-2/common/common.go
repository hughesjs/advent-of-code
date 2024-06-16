package common

type SeedMapEntry struct {
	SourceRange    SeedRange
	DestRangeStart int64
}

type SeedRange struct {
	Start int64
	End   int64
}

type SeedMap struct {
	Entries        []SeedMapEntry
	MapDescription string
}
