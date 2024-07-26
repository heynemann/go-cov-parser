package gocovparser

import "golang.org/x/tools/cover"

// ParseGroup to group coverage data by.
type ParseGroup struct {
	// Name of the parse group. Used to retrieve your parse data after grouping.
	Name string

	// KeyFunc that returns the grouping key to use based on the coverage line.
	KeyFunc func(string) string
}

// ParseGroupResult represents results of a Group Coverage operation.
type ParseGroupResult map[string]map[string]float64

// OverallCoverageBreakdown is an overall summary of code coverage.
type OverallCoverageBreakdown struct {
	TotalCoveredLines int
	TotalLines        int
	PercentByLines    float64

	TotalCoveredStatements int
	TotalStatements        int
	PercentByStatements    float64
}

// Filter interface for filtering coverage by.
type Filter interface {
	FilterCoverage(Coverage) bool
}

// Coverage line in a coverage.out file.
type Coverage struct {
	FileName string
	Blocks   []cover.ProfileBlock

	Host  string
	Owner string
	Repo  string
	Path  string
}
