#!/bin/bash
go clean -testcache
go test ./... \
    -timeout 1s \
    -v  \
    -p 1 \
    -cpu 1 \
    -failfast \
    -coverprofile=coverage.out
go tool cover -func=coverage.out
