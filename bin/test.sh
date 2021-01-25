#!/usr/bin/env bash

set -euo pipefail

docker-compose build test
docker-compose run --rm test
