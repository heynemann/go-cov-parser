package gocovparser_test

//revive:disable:add-constant

import (
	"testing"

	"github.com/heynemann/go-cov-parser/gocovparser"
	"github.com/stretchr/testify/assert"
)

func TestCanGroupCoverageData(t *testing.T) {
	const (
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
				groups: []gocovparser.ParseGroup{
					gocovparser.PackageParseGroup,
					gocovparser.FileParseGroup,
					gocovparser.TotalParseGroup,
				},
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result gocovparser.ParseGroupResult) {
				t.Helper()

				assert.Contains(t, result, "package")
				assert.Contains(t, result["package"], "github.com/heynemann/go-cov-parser/gocovparser")
				assert.EqualValues(t, result["package"]["github.com/heynemann/go-cov-parser/gocovparser"], expectedCov)

				assert.Contains(t, result, "file")
				assert.Contains(t, result["file"], "github.com/heynemann/go-cov-parser/gocovparser/core.go")
				assert.EqualValues(t, result["file"]["github.com/heynemann/go-cov-parser/gocovparser/core.go"], expectedCov)

				assert.Contains(t, result, "total")
				assert.Contains(t, result["total"], "total")
				assert.EqualValues(
					t,
					expectedCov,
					result["total"]["total"],
					result,
				)
			},
		},

		{
			name: "Can parse coverage data by package, file and total for another coverage file format",
			args: args{
				coverageData: CoverageFixture2(t),
				groups: []gocovparser.ParseGroup{
					gocovparser.PackageParseGroup,
					gocovparser.FileParseGroup,
					gocovparser.TotalParseGroup,
				},
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result gocovparser.ParseGroupResult) {
				t.Helper()

				assert.Contains(t, result, "package")
				assert.Contains(t, result["package"], "github.cbhq.net/engineering/mongofle")
				assert.EqualValues(t, 0.6918604651162791, result["package"]["github.cbhq.net/engineering/mongofle"], result)

				assert.Contains(t, result, "file")
				assert.Contains(t, result["file"], "github.cbhq.net/engineering/mongofle/mongo_encrypter.go")
				assert.EqualValues(
					t,
					0.8282208588957055,
					result["file"]["github.cbhq.net/engineering/mongofle/mongo_encrypter.go"],
					result,
				)

				assert.Contains(t, result, "total")
				assert.Contains(t, result["total"], "total")
				assert.EqualValues(t, 0.6918604651162791, result["total"]["total"], result)
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
			got, err := gocovparser.GroupCoverage(items, testcase.args.groups...)

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
