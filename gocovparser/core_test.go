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
				assert.EqualValues(t, expectedCov, result["package"]["github.com/heynemann/go-cov-parser/gocovparser"])

				assert.Contains(t, result, "file")
				assert.Contains(t, result["file"], "github.com/heynemann/go-cov-parser/gocovparser/core.go")
				assert.EqualValues(t, expectedCov, result["file"]["github.com/heynemann/go-cov-parser/gocovparser/core.go"])

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
			name: "Can parse coverage data by package, file and total for file without repo",
			args: args{
				coverageData: CoverageFixture5(t),
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
				assert.Contains(t, result["package"], "go.uber.org/zap")
				assert.EqualValues(t, 1, result["package"]["go.uber.org/zap"])

				assert.Contains(t, result, "file")
				assert.Contains(t, result["file"], "go.uber.org/zap/writer.go")
				assert.EqualValues(t, 1, result["file"]["go.uber.org/zap/writer.go"])

				assert.Contains(t, result, "total")
				assert.Contains(t, result["total"], "total")
				assert.EqualValues(
					t,
					1,
					result["total"]["total"],
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

		{
			name: "Can parse coverage data by package, file and total for another fixture type",
			args: args{
				coverageData: CoverageFixture7(t),
				groups: []gocovparser.ParseGroup{
					gocovparser.PackageParseGroup,
					gocovparser.FileParseGroup,
					gocovparser.TotalParseGroup,
				},
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result gocovparser.ParseGroupResult) {
				t.Helper()

				pkg := "github.cbhq.net/risk/data-tracker-backend/internal/serde"
				assert.Contains(t, result, "package")
				assert.Contains(t, result["package"], pkg)
				assert.EqualValues(t, 0.9545454545454546, result["package"][pkg], result)

				assert.Contains(t, result, "total")
				assert.Contains(t, result["total"], "total")
				assert.EqualValues(t, 0.7737819025522041, result["total"]["total"], result)
			},
		},

		{
			name: "Can parse coverage data by package, file in atomic mode",
			args: args{
				coverageData: CoverageFixture6(t),
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

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/cmd/server")
				assert.EqualValues(t, 0.1377952755905512, result["package"]["github.cbhq.net/some_repo/internal/cmd/server"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/config")
				assert.EqualValues(t, 0.4294117647058823, result["package"]["github.cbhq.net/some_repo/internal/config"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/acl")
				assert.EqualValues(t, 0.9594594594594594, result["package"]["github.cbhq.net/some_repo/internal/acl"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/middleware/desync")
				assert.EqualValues(t, 0.5, result["package"]["github.cbhq.net/some_repo/internal/middleware/desync"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/middleware/observability")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/middleware/observability"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/middleware/revoker")
				assert.EqualValues(t, 0.9, result["package"]["github.cbhq.net/some_repo/internal/middleware/revoker"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/middleware/throttler")
				assert.EqualValues(
					t,
					0.9333333333333333,
					result["package"]["github.cbhq.net/some_repo/internal/middleware/throttler"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/cache")
				assert.EqualValues(t, 0.8055555555555556, result["package"]["github.cbhq.net/some_repo/internal/pkg/cache"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/mcencryption")
				assert.EqualValues(
					t,
					0.6363636363636364,
					result["package"]["github.cbhq.net/some_repo/internal/pkg/mcencryption"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/pkmanager")
				assert.EqualValues(t, 0, result["package"]["github.cbhq.net/some_repo/internal/pkg/pkmanager"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/protojsonwrapper")
				assert.EqualValues(
					t,
					0.8571428571428571,
					result["package"]["github.cbhq.net/some_repo/internal/pkg/protojsonwrapper"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/resources")
				assert.EqualValues(t, 0.9647058823529412, result["package"]["github.cbhq.net/some_repo/internal/pkg/resources"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/safeconc")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/pkg/safeconc"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/service/v1")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/pkg/service/v1"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/storage")
				assert.EqualValues(t, 0.7192982456140351, result["package"]["github.cbhq.net/some_repo/internal/pkg/storage"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/storage/bsonbuilder")
				assert.EqualValues(
					t,
					0.9565217391304348,
					result["package"]["github.cbhq.net/some_repo/internal/pkg/storage/bsonbuilder"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/storage/composer")
				assert.EqualValues(
					t,
					0.9153094462540716,
					result["package"]["github.cbhq.net/some_repo/internal/pkg/storage/composer"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/storage/mongo")
				assert.EqualValues(
					t,
					0.8306342780026991,
					result["package"]["github.cbhq.net/some_repo/internal/pkg/storage/mongo"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/storage/paranoia")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/pkg/storage/paranoia"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/userscontext")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/pkg/userscontext"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/validators")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/pkg/validators"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/pkg/xtypes")
				assert.EqualValues(t, 1, result["package"]["github.cbhq.net/some_repo/internal/pkg/xtypes"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service1")
				assert.EqualValues(t, 0.824468085106383, result["package"]["github.cbhq.net/some_repo/internal/service/service1"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service2")
				assert.EqualValues(
					t,
					0.7714285714285715,
					result["package"]["github.cbhq.net/some_repo/internal/service/service2"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service3")
				assert.EqualValues(
					t,
					0.8431372549019608,
					result["package"]["github.cbhq.net/some_repo/internal/service/service3"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service4")
				assert.EqualValues(
					t,
					0.8472222222222222,
					result["package"]["github.cbhq.net/some_repo/internal/service/service4"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service5")
				assert.EqualValues(
					t,
					0.8271028037383178,
					result["package"]["github.cbhq.net/some_repo/internal/service/service5"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service6")
				assert.EqualValues(
					t,
					0.8382352941176471,
					result["package"]["github.cbhq.net/some_repo/internal/service/service6"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service7")
				assert.EqualValues(
					t,
					0.8571428571428571,
					result["package"]["github.cbhq.net/some_repo/internal/service/service7"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service8")
				assert.EqualValues(
					t,
					0.6166666666666667,
					result["package"]["github.cbhq.net/some_repo/internal/service/service8"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service9")
				assert.EqualValues(
					t,
					0.8498293515358362,
					result["package"]["github.cbhq.net/some_repo/internal/service/service9"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service10")
				assert.EqualValues(
					t,
					0.8433734939759037,
					result["package"]["github.cbhq.net/some_repo/internal/service/service10"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service/service11")
				assert.EqualValues(
					t,
					0.8478260869565217,
					result["package"]["github.cbhq.net/some_repo/internal/service/service/service11"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service/service12")
				assert.EqualValues(
					t,
					0.8367346938775511,
					result["package"]["github.cbhq.net/some_repo/internal/service/service/service12"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service/service13")
				assert.EqualValues(t, 0.859375, result["package"]["github.cbhq.net/some_repo/internal/service/service/service13"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service/service14")
				assert.EqualValues(
					t,
					0.8481012658227848,
					result["package"]["github.cbhq.net/some_repo/internal/service/service/service14"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/service/service15")
				assert.EqualValues(t, 0.84375, result["package"]["github.cbhq.net/some_repo/internal/service/service/service15"])

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/internal/service/v1/utils")
				assert.EqualValues(
					t,
					0.9302325581395349,
					result["package"]["github.cbhq.net/some_repo/internal/service/v1/utils"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/lib/admin_flags/pkg/afsdk")
				assert.EqualValues(
					t,
					0.8571428571428571,
					result["package"]["github.cbhq.net/some_repo/lib/admin_flags/pkg/afsdk"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/lib/admin_flags/pkg/protogen")
				assert.EqualValues(
					t,
					0.8367346938775511,
					result["package"]["github.cbhq.net/some_repo/lib/admin_flags/pkg/protogen"],
				)

				assert.Contains(t, result["package"], "github.cbhq.net/some_repo/lib/admin_flags/pkg/validator")
				assert.EqualValues(
					t,
					0.9705882352941176,
					result["package"]["github.cbhq.net/some_repo/lib/admin_flags/pkg/validator"],
				)

				assert.Contains(t, result["file"], "github.cbhq.net/some_repo/internal/cmd/server/server.go")
				assert.EqualValues(
					t,
					0.22900763358778625,
					result["file"]["github.cbhq.net/some_repo/internal/cmd/server/server.go"],
				)
				assert.Contains(t, result["file"], "github.cbhq.net/some_repo/internal/config/config.go")
				assert.EqualValues(t, 0.8490566037735849, result["file"]["github.cbhq.net/some_repo/internal/config/config.go"])

				assert.Contains(t, result, "total")
				assert.Contains(t, result["total"], "total")
				assert.EqualValues(
					t,
					0.7949806237313157,
					result["total"]["total"],
				)
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
			defer func() {
				if t.Failed() {
					t.Logf("items: %#v\n", items)
				}
			}()

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

				assert.Len(t, result, 1)
				assert.Len(t, result[0].Blocks, 43)
			},
		},
		{
			name: "Can parse coverage data by package, file and total if coverage is empty",
			args: args{
				coverageData: EmptyFixture(t),
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result []gocovparser.Coverage) {
				t.Helper()

				assert.Len(t, result, 0)
			},
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

func TestCanGetTotalCoverageBreakdown(t *testing.T) {
	const (
		expectedTotalCoveredLines      = 112
		expectedTotalLines             = 134
		expectedPercentByLines         = 0.835820895522388
		expectedTotalCoveredStatements = 60
		expectedTotalStatements        = 68
		expectedPercentByStatement     = 0.8823529411764706
	)

	type args struct {
		coverageData string
	}

	tests := []struct {
		name         string
		args         args
		expectedErr  error
		expectedData func(t *testing.T, result gocovparser.OverallCoverageBreakdown)
	}{
		{
			name: "Can parse coverage data by package, file and total",
			args: args{
				coverageData: CoverageFixture(t),
			},
			expectedErr: nil,
			expectedData: func(t *testing.T, result gocovparser.OverallCoverageBreakdown) {
				t.Helper()

				assert.EqualValues(t, expectedTotalCoveredLines, result.TotalCoveredLines)
				assert.EqualValues(t, expectedTotalLines, result.TotalLines)

				assert.EqualValues(t, expectedTotalCoveredStatements, result.TotalCoveredStatements)
				assert.EqualValues(t, expectedTotalStatements, result.TotalStatements)

				assert.EqualValues(t, expectedPercentByLines, result.PercentByLines)
				assert.EqualValues(t, expectedPercentByStatement, result.PercentByStatements)
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
			defer func() {
				if t.Failed() {
					t.Logf("items: %#v\n", items)
				}
			}()

			// ACT
			got, err := gocovparser.GetTotalCoverageBreakdown(items)

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
