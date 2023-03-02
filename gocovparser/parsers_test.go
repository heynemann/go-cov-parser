package gocovparser_test

//revive:disable:add-constant

import (
	"fmt"
	"math"
	"testing"

	"github.com/heynemann/go-cov-parser/gocovparser"
	"github.com/stretchr/testify/require"
)

func getCov(t *testing.T, got gocovparser.ParseGroupResult, parser, key string) int {
	t.Helper()

	require.Contains(t, got, parser)

	parserData := got[parser]

	require.Contains(t, parserData, key)

	coverage := parserData[key]

	return int(
		math.Round(coverage * 10000.0),
	)
}

func TestFileParser(t *testing.T) {
	data1 := CoverageFixture(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.FileParseGroup)

	// ASSERT
	require.NoError(t, err)

	path := "github.com/heynemann/go-cov-parser/gocovparser/core.go"
	require.Equal(t, 8824, getCov(t, got, "file", path))
}

func TestFileParser2(t *testing.T) {
	data1 := CoverageFixture2(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.FileParseGroup)

	// ASSERT
	require.NoError(t, err)

	path := "github.cbhq.net/engineering/mongofle/mongo_encrypter.go"
	require.Equal(t, 8282, getCov(t, got, "file", path))
}

func TestPackageParser(t *testing.T) {
	data1 := CoverageFixture(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.PackageParseGroup)

	// ASSERT
	require.NoError(t, err)

	path := "github.com/heynemann/go-cov-parser/gocovparser"
	require.Equal(t, 8824, getCov(t, got, "package", path))
}

func TestPackageParser2(t *testing.T) {
	data1 := CoverageFixture2(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.PackageParseGroup)

	// ASSERT
	require.NoError(t, err)

	path := "github.cbhq.net/engineering/mongofle"
	require.Equal(t, 6919, getCov(t, got, "package", path))
}

func TestPackageParser3(t *testing.T) {
	data1 := CoverageFixture3(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.PackageParseGroup)

	// ASSERT
	require.NoError(t, err)

	pkg := func(item string) string { return fmt.Sprintf("github.cbhq.net/engineering/docstore/%s", item) }

	require.Equal(t, 7318, getCov(t, got, "package", pkg("bsonpb")))
	require.Equal(t, 4545, getCov(t, got, "package", pkg("testhelpers")))
	require.Equal(t, 1938, getCov(t, got, "package", pkg("sdk")))
	require.Equal(t, 9222, getCov(t, got, "package", pkg("modules/mongo")))
}

func TestTotalParser(t *testing.T) {
	data1 := CoverageFixture(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.TotalParseGroup)

	// ASSERT
	require.NoError(t, err)

	require.Equal(t, 8824, getCov(t, got, "total", "total"))
}

func TestTotalParser2(t *testing.T) {
	data1 := CoverageFixture2(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.TotalParseGroup)

	// ASSERT
	require.NoError(t, err)

	coverage := got["total"]["total"]
	require.Equal(t, float64(6919), math.Round(coverage*10000.0))
}

func TestTotalParser3(t *testing.T) {
	data1 := CoverageFixture3(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.TotalParseGroup)

	// ASSERT
	require.NoError(t, err)

	coverage := got["total"]["total"]
	require.Equal(t, float64(7497), math.Round(coverage*10000.0))
}

func TestTotalParser4(t *testing.T) {
	data1 := CoverageFixture4(t)

	items, err := gocovparser.Parse(data1)
	require.NoError(t, err)

	// ACT
	got, err := gocovparser.GroupCoverage(items, gocovparser.TotalParseGroup)

	// ASSERT
	require.NoError(t, err)

	coverage := got["total"]["total"]
	require.Equal(t, float64(195), math.Round(coverage*10000.0))
}
