package gocovparser_test

//revive:disable:add-constant

import (
	"strings"
	"testing"

	"github.com/heynemann/go-cov-parser/gocovparser"
	"github.com/stretchr/testify/assert"
)

func TestCanGroupCoverageData(t *testing.T) {
	var (
		packageParseGroup = gocovparser.ParseGroup{
			Name: "package",
			KeyFunc: func(line string) string {
				line = strings.ReplaceAll(line, repoName+"/", "")
				parts := strings.Split(line, "/")

				return strings.Join(parts[:len(parts)-1], "/")
			},
		}
		fileParseGroup = gocovparser.ParseGroup{
			Name: "file",
			KeyFunc: func(line string) string {
				line = strings.ReplaceAll(line, repoName+"/", "")

				return strings.ReplaceAll(strings.Split(line, ":")[0], "github.com/heynemann/gocovparser", "")
			},
		}
		totalParseGroup = gocovparser.ParseGroup{
			Name:    "total",
			KeyFunc: func(line string) string { return "total" },
		}

		expectedCov = 0.8823529411764706
	)

	type args struct {
		coverageData string
		groups       []gocovparser.ParseGroup
	}

	tests := []struct {
		name         string
		args         args
		expectedErr  error
		expectedData func(t *testing.T, result gocovparser.ParseGroupResult)
	}{
		{
			name: "Can parse coverage data by package, file and total",
			args: args{
				coverageData: CoverageFixture(t),
				groups:       []gocovparser.ParseGroup{packageParseGroup, fileParseGroup, totalParseGroup},
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result gocovparser.ParseGroupResult) {
				t.Helper()

				assert.Contains(t, result, "package")
				assert.Contains(t, result["package"], "gocovparser")
				assert.EqualValues(t, result["package"]["gocovparser"], expectedCov)

				assert.Contains(t, result, "file")
				assert.Contains(t, result["file"], "gocovparser/core.go")
				assert.EqualValues(t, result["file"]["gocovparser/core.go"], expectedCov)

				assert.Contains(t, result, "total")
				assert.Contains(t, result["total"], "total")
				assert.EqualValues(t, result["total"]["total"], expectedCov)
			},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			// ARRANGE
			items, err := gocovparser.Parse(testcase.args.coverageData)
			if !assert.NoError(t, err) {
				return
			}

			// ACT
			got, err := gocovparser.GroupCoverage(items, testcase.args.groups)

			// ASSERT
			if testcase.expectedErr != nil {
				if assert.Error(t, err) {
					assert.ErrorIs(t, err, testcase.expectedErr)
				}

				return
			}

			if assert.NoError(t, err) {
				assert.NotNil(t, got)
				testcase.expectedData(t, got)
			}
		})
	}
}

func TestCanParseCoverageData(t *testing.T) {
	type args struct {
		coverageData string
	}

	tests := []struct {
		name         string
		args         args
		expectedErr  error
		expectedData func(*testing.T, []gocovparser.Coverage)
	}{
		{
			name: "Can parse coverage data by package, file and total",
			args: args{
				coverageData: CoverageFixture(t),
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result []gocovparser.Coverage) {
				t.Helper()

				assert.Len(t, result, 43)
			},
		},
		{
			name: "Fail if empty data",
			args: args{
				coverageData: "",
			},
			expectedErr: gocovparser.ErrInvalidCoverageData,
		},
		{
			name: "Fail if invalid coverage line",
			args: args{
				coverageData: `mode: set
github.com/heynemann/go-cov-parser/gocovparser/core.go:invalid`,
			},
			expectedErr: gocovparser.ErrInvalidCoverageData,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			// ACT
			got, err := gocovparser.Parse(testcase.args.coverageData)

			// ASSERT
			if testcase.expectedErr != nil {
				if assert.Error(t, err) {
					assert.ErrorIs(t, err, testcase.expectedErr)
				}

				return
			}

			if assert.NoError(t, err) {
				assert.NotNil(t, got)
				testcase.expectedData(t, got)
			}
		})
	}
}
