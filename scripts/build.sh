#!/usr/bin/env bash

echo "==> Removing old directory..."
rm -f bin/*
mkdir -p bin/

echo "==> Building..."
go build  -tags=jsoniter -o ./bin ./...

echo "==> Results:"
ls -hl bin/