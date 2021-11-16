#!/bin/bash

set -e

find . -name '*.go' -not -wholename './vendor/*' |
while read -r file
do
  golines -m 140 -w "$file"
  gci -w "$file"
  gofumpt -w "$file"
done
