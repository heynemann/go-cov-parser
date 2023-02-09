.PHONY: all test clean

all: test clean

clean:
	@rm -rf coverage.out
	@rm -rf dist

test:
	@gotestsum --format testname -- -coverprofile=coverage.out ./...

watch:
	@gotestsum --format testname --watch -- -coverprofile=coverage.out ./...

release:
	@goreleaser release --clean
