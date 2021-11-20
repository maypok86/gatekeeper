#!/bin/bash

set -e

docker-compose -f deployments/docker-compose.test.yml up --build -d
test_status_code=0
docker-compose -f deployments/docker-compose.test.yml run integration_tests go test ./tests/integration || test_status_code=$?
docker-compose -f deployments/docker-compose.test.yml logs > integration.log
docker-compose -f deployments/docker-compose.test.yml down
exit $test_status_code
