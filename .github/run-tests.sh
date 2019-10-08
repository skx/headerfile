#!/bin/bash

# Install tools to test our code-quality.
go get -u golang.org/x/lint/golint
go get -u honnef.co/go/tools/cmd/staticcheck

# Init the modules
go mod init

# Run the static-check tool - we ignore errors in goserver/static.go
t=$(mktemp)
staticcheck -checks all ./... > $t
if [ -s $t ]; then
    echo "Found errors via 'staticcheck'"
    cat $t
    rm $t
    exit 1
fi
rm $t

# At this point failures cause aborts
set -e

# Run the linter
echo "Launching 'golint' check .."
golint -set_exit_status ./...
echo "Completed 'golint' check .."

# Run the vet-check
echo "Launching 'go vet' check .."
go vet ./...
echo "Completed 'go vet' check .."

# Run our package tests
go test ./...

#
# Test coverage not being 100% is a bug.
#
coverage=$(go test -coverprofile=tmp | grep coverage | awk '{print $2}')

if [ "${coverage}" == "100.0%" ]; then
    echo "100% test-coverage.  Good job"
else
    echo "Coverage is ${coverage} not 100%"
    exit 1
fi
