package gocovparser

import "errors"

// ErrInvalidCoverageData happens when the data passed to gocovparser is either blank or not a coverage.out file content.
var ErrInvalidCoverageData = errors.New("invalid coverage data - unable to parse")
