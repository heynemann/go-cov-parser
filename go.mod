module github.com/heynemann/go-cov-parser

go 1.19

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	golang.org/x/tools v0.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.3.4 // Contains retractions only.
	v0.3.0 // Version with problems.
)
