package gocovparser_test

import "testing"

const repoName = "github.com/heynemann/go-cov-parser"

// CoverageFixture is a mock coverage.out contents.
func CoverageFixture(t *testing.T) string {
	t.Helper()

	return `
mode: set
github.com/heynemann/go-cov-parser/gocovparser/core.go:38.53,42.2 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:45.60,47.20 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:51.2,53.30 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:110.2,110.22 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:47.20,49.3 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:53.30,55.22 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:59.3,65.17 6 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:69.3,70.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:74.3,75.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:79.3,80.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:84.3,85.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:89.3,107.4 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:55.22,57.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:65.17,67.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:70.17,72.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:75.17,77.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:80.17,82.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:85.17,87.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:113.43,116.29 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:130.2,130.15 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:116.29,119.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:123.3,123.39 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:127.3,127.32 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:119.17,120.12 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:123.39,124.12 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:133.40,135.16 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:139.2,139.25 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:135.16,137.3 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:143.85,149.31 4 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:186.2,186.20 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:149.31,150.45 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:154.3,154.49 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:158.3,158.46 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:162.3,162.29 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:172.3,172.43 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:150.45,152.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:154.49,156.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:158.46,160.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:162.29,167.21 3 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:167.21,169.5 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:172.43,176.18 3 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:182.4,182.59 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:176.18,179.13 2 0
`
}
