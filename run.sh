#!/usr/bin/env bash


set -o errexit
set -o pipefail
set -o nounset

go run \
  config.go \
  convertToRDAP.go \
  main.go \
  rdapHandler.go \
  staticHandler.go \
  status.go \
  whoisLookup.go 
