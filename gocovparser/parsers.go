package gocovparser

import (
	"strings"
)

// FileParseGroup returns each line as it's own key.
var FileParseGroup = ParseGroup{
	Name: "file",
	KeyFunc: func(filename string) string {
		return filename
	},
}

// PackageParseGroup parses the package from the coverage data and uses that as identifier.
var PackageParseGroup = ParseGroup{
	Name: "package",
	KeyFunc: func(filename string) string {
		parts := strings.Split(filename, "/")

		return strings.Join(parts[:len(parts)-1], "/")
	},
}

// TotalParseGroup uses 'total' as key to group all coverage data into a total.
var TotalParseGroup = ParseGroup{
	Name: "total",
	KeyFunc: func(string) string {
		return "total"
	},
}
