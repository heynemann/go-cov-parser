package gocovparser

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/cover"
)

const (
	hostPosition  = 1
	ownerPosition = 2
	repoPosition  = 3
	pathPosition  = 4
)

var parseLineRegex = regexp.MustCompile(
	`(?P<host>[^\/]*)\/` + // github.com
		`(?P<owner>[^\/]*)\/` + // heynemann
		`(?:(?P<repo>[^\/]*)\/)?` + // gocovparser
		`(?P<path>.*)`, // gocovparser/core.go
)

// Parse a coverage result file contents from go tests.
func Parse(coverageData string) ([]Coverage, error) {
	// Remove empty blank lines
	coverageData = strings.TrimSpace(coverageData)

	profiles, err := cover.ParseProfilesFromReader(strings.NewReader(coverageData))
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidCoverageData, err.Error())
	}

	if len(profiles) == 0 {
		return []Coverage{}, nil
	}

	coverage := make([]Coverage, 0, len(profiles))

	for _, profile := range profiles {
		match := parseLineRegex.FindStringSubmatch(profile.FileName)
		if len(match) == 0 {
			return nil, errors.Wrapf(ErrInvalidCoverageData, "invalid coverage file name %q", profile.FileName)
		}

		host := match[hostPosition]
		owner := match[ownerPosition]
		repo := match[repoPosition]
		path := match[pathPosition]

		coverage = append(coverage, Coverage{
			FileName: profile.FileName,
			Host:     host,
			Owner:    owner,
			Repo:     repo,
			Path:     path,
			Blocks:   profile.Blocks,
		})
	}

	return coverage, nil
}

// GroupCoverage in the specified groups.
func GroupCoverage(items []Coverage, groups ...ParseGroup) (ParseGroupResult, error) {
	rawCov, err := GroupCoverageRaw(items, groups...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]float64)

	for group, values := range rawCov {
		if _, found := result[group]; !found {
			result[group] = make(map[string]float64)
		}

		for key, raw := range values {
			if raw.NumStatements == 0 {
				result[group][key] = 0.0
			} else {
				result[group][key] = float64(raw.CoveredStatements) / float64(raw.NumStatements)
			}
		}
	}

	return result, nil
}

// GroupCoverageRaw reports raw lines and statements instead of percentages
func GroupCoverageRaw(items []Coverage, groups ...ParseGroup) (ParseGroupRawResult, error) {
	result := make(map[string]map[string]*RawCoverage)

	for _, group := range groups {
		if _, found := result[group.Name]; !found {
			result[group.Name] = make(map[string]*RawCoverage)
		}

		for _, cov := range items {
			key := group.KeyFunc(cov.FileName)
			if _, found := result[group.Name][key]; !found {
				result[group.Name][key] = &RawCoverage{}
			}

			for _, b := range cov.Blocks {
				nlines := b.EndLine - b.StartLine

				result[group.Name][key].NumLines += nlines
				result[group.Name][key].NumStatements += b.NumStmt

				if b.Count > 0 { // Count is the number of times that block was covered
					result[group.Name][key].CoveredLines += nlines
					result[group.Name][key].CoveredStatements += b.NumStmt
				}
			}
		}
	}

	return result, nil
}
