#!/usr/bin/env bash

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(gofmt -l `find . -name '*.go' | grep -v vendor`)
if [[ -n ${gofmt_files} ]]; then
    echo "Use the command \`make fmt\` to reformat this code:"
    echo "${gofmt_files}"
    exit 1
fi

exit 0
