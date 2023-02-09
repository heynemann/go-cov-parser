.PHONY: all test clean

all: test clean

clean:
	@rm -rf coverage.out

test:
	@gotestsum --format testname -- -coverprofile=coverage.out ./...

watch:
	@gotestsum --format testname --watch -- -coverprofile=coverage.out ./...
