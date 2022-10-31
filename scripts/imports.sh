#!/bin/bash

# Format
#go fmt `go list ./... | grep -v "/vendor/" | grep -v "/gen/"`
#gofumpt -l -w `go list ./... | grep -v "/vendor/" | grep -v "/gen/"`
gofumpt -l -w .

# Find target files
gofiles=$(find . -name "*.go" | grep -v "/gen/")

# Remove blank line in import section
#FIXME: somehow blank line between functions is gone in only _test.go files
for gofile in $gofiles; do
    echo $gofile
    sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile > tmp
    mv tmp $gofile
done

# goimports
goimports -local github.com/hiromaily/,resume/gen/ -w `goimports -local github.com/hiromaily/,resume/gen/ -l ./ | grep -v "/gen/"`
