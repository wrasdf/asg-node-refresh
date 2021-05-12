#!/usr/bin/env bash

set -euo pipefail

docker-compose build sh
docker-compose run --rm sh -c "go test ./tests"
