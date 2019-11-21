#!/usr/bin/env bash

@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
    go get -u golang.org/x/tools/cmd/cover; \
fi
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
rm coverage.out