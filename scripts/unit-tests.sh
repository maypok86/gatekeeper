#!/bin/bash

set -e

echo 'mode: atomic' > coverage.txt
go test -covermode=atomic -coverprofile=coverage.txt -v -race ./cmd/... ./internal/... ./pkg/...
