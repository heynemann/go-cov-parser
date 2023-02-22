package gocovparser

// ParseGroup to group coverage data by.
type ParseGroup struct {
	// Name of the parse group. Used to retrieve your parse data after grouping.
	Name string

	// KeyFunc that returns the grouping key to use based on the coverage line.
	KeyFunc func(string) string
}

// ParseGroupResult represents results of a Group Coverage operation.
type ParseGroupResult map[string]map[string]float64

// Filter interface for filtering coverage by.
type Filter interface {
	FilterCoverage(Coverage) bool
}

// Coverage line in a coverage.out file.
type Coverage struct {
	Line       string
	Host       string
	Owner      string
	Repo       string
	Path       string
	Start      CoveragePosition
	End        CoveragePosition
	Statements int
	IsCovered  bool
}

// CoveragePosition indicates the start or end of a given coverage information.
type CoveragePosition struct {
	LineNo int
	Column int
}
