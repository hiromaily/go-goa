#!/bin/bash

# Format
#go fmt `go list ./... | grep -v "/vendor/" | grep -v "/gen/"`
#gofumpt -l -w .
gofumpt -l -w `go list ./... | grep -v "/vendor/" | grep -v "/gen/"`

# Find target files
gofiles=$(find . -name "*.go" | grep -v "/gen/")

# Remove blank line in import section
for gofile in $gofiles; do
    echo $gofile
    sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile > tmp
    mv tmp $gofile
done

goimports -local github.com/hiromaily/,resume/gen/ -w `goimports -local github.com/hiromaily/,resume/gen/ -l ./ | grep -v "/gen/"`
