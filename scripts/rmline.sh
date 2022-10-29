#!/bin/bash

# Find target files
gofiles=$(find . -name "*.go" | grep -v "/gen/")

# Remove blank line in import section
for gofile in $gofiles; do
    echo $gofile
    sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile > tmp
    mv tmp $gofile
done
