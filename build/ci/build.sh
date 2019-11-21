#!/usr/bin/env bash

docker build -f ./build/package/Dockerfile -t stoksc/hello .

docker push stoksc/hello