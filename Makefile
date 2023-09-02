# The -failfast flag tells go test to stop running tests at the first failure.
# The -v flag makes the test output verbose, showing all test results, not just the failures.
# The -timeout 30m flag sets a timeout of 30 minutes for the tests, after which they will be forcefully stopped.
.PHONY: test-unit
test-unit:
	cd server && go test ./ -failfast -v -timeout 30m
	cd coinbase && go test ./ -failfast -v timeout 30m