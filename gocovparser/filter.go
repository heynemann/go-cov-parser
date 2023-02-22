package gocovparser

import (
	"fmt"
	"regexp"
	"strings"
)

type packageExcludeFilter struct {
	packageName string
}

var _ Filter = (*packageExcludeFilter)(nil)

// PackageExcludeFilter excludes any coverage that matches the packageName specified.
func PackageExcludeFilter(packageName string) Filter {
	return &packageExcludeFilter{
		packageName: packageName,
	}
}

func (f *packageExcludeFilter) FilterCoverage(cov Coverage) bool {
	fullAddr := fmt.Sprintf("%s/%s/%s/%s", cov.Host, cov.Owner, cov.Repo, cov.Path)

	return !strings.HasPrefix(fullAddr, f.packageName)
}

type fileExcludeFilter struct {
	fileglob *regexp.Regexp
}

var _ Filter = (*fileExcludeFilter)(nil)

// FileExcludeFilter excludes any coverage that has a filename matching the specified regex.
func FileExcludeFilter(fileMatchRE string) Filter {
	fileglobRE := regexp.MustCompile(fileMatchRE)

	return &fileExcludeFilter{
		fileglob: fileglobRE,
	}
}

func (f *fileExcludeFilter) FilterCoverage(cov Coverage) bool {
	return !f.fileglob.MatchString(cov.Path)
}

// FilterCoverage using the specified filters.
func FilterCoverage(items []Coverage, filters ...Filter) ([]Coverage, error) {
	result := []Coverage{}

	for _, item := range items {
		keep := true

		for _, filter := range filters {
			if filter.FilterCoverage(item) {
				continue
			}

			keep = false

			break
		}

		if keep {
			result = append(result, item)
		}
	}

	return result, nil
}
