#!/usr/bin/env bash


set -o errexit
set -o pipefail
set -o nounset

go run main.go \
  convertToRDAP.go \
  staticHandler.go \
  whoisLookup.go \
  rdapHandler.go \
  config.go
