package gocovparser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	hostPosition       = 1
	ownerPosition      = 2
	repoPosition       = 3
	pathPosition       = 4
	lineno1Position    = 5
	column1Position    = 6
	lineno2Position    = 7
	column2Position    = 8
	statementsPosition = 9
	isCoveredPosition  = 10
	int64Bits          = 64
	intBase            = 10
)

var parseLineRegex = regexp.MustCompile(
	`(?P<host>[^\/]*)\/` + // github.com
		`(?P<owner>[^\/]*)\/` + // heynemann
		`(?P<repo>[^\/]*)\/` + // gocovparser
		`(?P<path>.*)\:` + // gocovparser/core.go
		`(?P<lineno1>\d+)[.](?P<column1>\d+)\,` + // 10.73
		`(?P<lineno2>\d+)[.](?P<column2>\d+)\s` + // 14.2
		`(?P<statements>\d+)\s` + // 2
		`(?P<iscov>\d+)`, // 1
)

// Parse a coverage result file contents from go tests.
func Parse(coverageData string) ([]Coverage, error) {
	lines := strings.Split(coverageData, "\n")

	return ParseLines(lines)
}

// ParseLines of a coverage file contents from go tests.
func ParseLines(coverageData []string) ([]Coverage, error) {
	data := filterEmpty(coverageData)
	if len(data) == 0 {
		return nil, errors.Wrap(ErrInvalidCoverageData, "empty coverage data")
	}

	coverage := make([]Coverage, len(data))

	for pos, line := range data {
		match := parseLineRegex.FindStringSubmatch(line)
		if len(match) == 0 {
			return nil, errors.Wrapf(ErrInvalidCoverageData, "invalid coverage line '%s'", line)
		}

		host := match[hostPosition]
		owner := match[ownerPosition]
		repo := match[repoPosition]
		path := match[pathPosition]

		lineno1, err := parseInt(match[lineno1Position])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid line number 1: '%s'", match[lineno1Position])
		}

		column1, err := parseInt(match[column1Position])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid column 1: '%s'", match[column1Position])
		}

		lineno2, err := parseInt(match[lineno2Position])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid line number 2: '%s'", match[lineno2Position])
		}

		column2, err := parseInt(match[column2Position])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid column 2: '%s'", match[column2Position])
		}

		statements, err := parseInt(match[statementsPosition])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid number of statements: '%s'", match[statementsPosition])
		}

		isCovered := match[isCoveredPosition] != "0"

		coverage[pos] = Coverage{
			Line:  line,
			Host:  host,
			Owner: owner,
			Repo:  repo,
			Path:  path,
			Start: CoveragePosition{
				LineNo: lineno1,
				Column: column1,
			},
			End: CoveragePosition{
				LineNo: lineno2,
				Column: column2,
			},
			Statements: statements,
			IsCovered:  isCovered,
		}
	}

	return coverage, nil
}

func filterEmpty(slice []string) []string {
	result := []string{}

	for _, item := range slice {
		line := strings.TrimSpace(item)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "mode:") {
			continue
		}

		result = append(result, item)
	}

	return result
}

func parseInt(val string) (int, error) {
	result, err := strconv.ParseInt(val, intBase, int64Bits)
	if err != nil {
		return -1, errors.Wrapf(err, "parsing int: '%s'", val)
	}

	return int(result), nil
}

// GroupCoverage in the specified groups.
func GroupCoverage(items []Coverage, groups ...ParseGroup) (ParseGroupResult, error) {
	result := map[string]map[string]float64{}

	statements := map[string]map[string]int{}
	covered := map[string]map[string]int{}

	for _, group := range groups {
		if _, found := result[group.Name]; !found {
			result[group.Name] = map[string]float64{}
		}

		if _, found := statements[group.Name]; !found {
			statements[group.Name] = map[string]int{}
		}

		if _, found := covered[group.Name]; !found {
			covered[group.Name] = map[string]int{}
		}

		for _, cov := range items {
			key := group.KeyFunc(cov.Line)

			statements[group.Name][key] += cov.Statements

			if cov.IsCovered {
				covered[group.Name][key] += cov.Statements
			}
		}

		for key := range statements[group.Name] {
			stmts := statements[group.Name][key]
			cov := covered[group.Name][key]

			if stmts == 0 {
				result[group.Name][key] = 0.0

				continue
			}

			result[group.Name][key] = float64(cov) / float64(stmts)
		}
	}

	return result, nil
}
