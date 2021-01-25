#!/usr/bin/env bash

# set -euo pipefail

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o ./bin/manager main.go

go run ./main.go
