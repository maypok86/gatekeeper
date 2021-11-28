#!/bin/bash

set -e

VERSION=$(git describe --tags --abbrev=0)
DATE=$(date -u +%Y-%m-%dT%H:%M:%S)
LDFLAGS="-X 'github.com/maypok86/gatekeeper/cmd.version=$VERSION' -X 'github.com/maypok86/gatekeeper/cmd.buildDate=$DATE'"

LDFLAGS=$LDFLAGS docker-compose -f deployments/docker-compose.yml --env-file .env up -d --build
docker image prune -f
