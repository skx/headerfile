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
echo "Launching linter .."
golint -set_exit_status ./...
echo "Completed linter .."

# Run the vet-check
echo "Launching shadowed-variable check .."
go vet ./...
echo "Completed shadowed-variable check .."

# Run golang tests
go test ./...
