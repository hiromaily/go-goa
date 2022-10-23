#!/bin/bash

gofiles=$(find . -name "*.go" | grep -v "/gen/")

for gofile in $gofiles; do
    echo $gofile
    sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile > tmp
    mv tmp $gofile
done

go fmt `go list ./... | grep -v "/vendor/"`
goimports -local github.com/hiromaily/,resume/gen/ -w `goimports -local github.com/hiromaily/,resume/gen/ -l ./ | grep -v "/gen/"`
