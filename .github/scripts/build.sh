#!/bin/bash
set -xe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./app/hotel-rates-api ./
tar -czvf rates-api.tar.gz ./app/hotel-rates-api ./appspec.yml