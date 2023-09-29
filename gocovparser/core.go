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
	result := make(map[string]map[string]float64)

	statements := make(map[string]map[string]int)
	covered := make(map[string]map[string]int)

	for _, group := range groups {
		if _, found := result[group.Name]; !found {
			result[group.Name] = make(map[string]float64)
		}

		if _, found := statements[group.Name]; !found {
			statements[group.Name] = make(map[string]int)
		}

		if _, found := covered[group.Name]; !found {
			covered[group.Name] = make(map[string]int)
		}

		for _, cov := range items {
			key := group.KeyFunc(cov.FileName)

			for _, b := range cov.Blocks {
				statements[group.Name][key] += b.NumStmt

				if b.Count > 0 { // is covered
					covered[group.Name][key] += b.NumStmt
				}
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
