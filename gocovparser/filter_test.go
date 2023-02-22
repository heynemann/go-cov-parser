package gocovparser_test

//revive:disable:add-constant

import (
	"testing"

	"github.com/heynemann/go-cov-parser/gocovparser"
	"github.com/stretchr/testify/require"
)

func TestPackageFilter(t *testing.T) {
	data1 := CoverageFixture(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.FilterCoverage(
		items,
		gocovparser.PackageExcludeFilter("github.com/heynemann/go-cov-parser/gocovparser"),
	)

	// ASSERT
	require.NoError(t, err)

	require.Len(t, got, 0)
}

func TestFileFilter(t *testing.T) {
	data1 := CoverageFixture2(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.FilterCoverage(
		items,
		gocovparser.FileExcludeFilter("default_mongo_encrypter\\.go"),
		gocovparser.FileExcludeFilter("key_provider\\.go"),
		gocovparser.FileExcludeFilter("mongo_encrypter\\.go"),
	)

	// ASSERT
	require.NoError(t, err)

	require.Len(t, got, 32)
}

// func TestPackageParser(t *testing.T) {
// data1 := CoverageFixture(t)

// items, err := gocovparser.Parse(data1)
// require.NoError(t, err)

// // ACT
// got, err := gocovparser.GroupCoverage(items, gocovparser.PackageParseGroup)

// // ASSERT
// require.NoError(t, err)

// path := "github.com/heynemann/go-cov-parser/gocovparser"
// require.Equal(t, 8824, getCov(t, got, "package", path))
// }

// func TestPackageParser2(t *testing.T) {
// data1 := CoverageFixture2(t)

// items, err := gocovparser.Parse(data1)
// require.NoError(t, err)

// // ACT
// got, err := gocovparser.GroupCoverage(items, gocovparser.PackageParseGroup)

// // ASSERT
// require.NoError(t, err)

// path := "github.cbhq.net/engineering/mongofle"
// require.Equal(t, 6919, getCov(t, got, "package", path))
// }

// func TestTotalParser(t *testing.T) {
// data1 := CoverageFixture(t)

// items, err := gocovparser.Parse(data1)
// require.NoError(t, err)

// // ACT
// got, err := gocovparser.GroupCoverage(items, gocovparser.TotalParseGroup)

// // ASSERT
// require.NoError(t, err)

// require.Equal(t, 8824, getCov(t, got, "total", "total"))
// }

// func TestTotalParser2(t *testing.T) {
// data1 := CoverageFixture2(t)

// items, err := gocovparser.Parse(data1)
// require.NoError(t, err)

// // ACT
// got, err := gocovparser.GroupCoverage(items, gocovparser.TotalParseGroup)

// // ASSERT
// require.NoError(t, err)

// coverage := got["total"]["total"]
// require.Equal(t, float64(6919), math.Round(coverage*10000.0))
// }
