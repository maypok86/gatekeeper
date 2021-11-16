#!/bin/bash

set -e

GIT_HASH=$(git log --format="%h" -n 1)
LDFLAGS="-X 'github.com/maypok86/gatekeeper/cmd.version=$GIT_HASH' -X 'github.com/maypok86/gatekeeper/cmd.buildDate=$(date -u +%Y-%m-%dT%H:%M:%S)'"

go build -v -o "$1" -ldflags "$LDFLAGS" .
